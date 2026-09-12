package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitResultRelation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = l2 - int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9<<(uint(int32(2))%32))))
	if v13 == int32(0) {
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+l2<<(uint(int32(2))%32)-int32(4))))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
		v27 = *(*int32)(unsafe.Add(mBase, _consts[46]))
		if v16 <= v27 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
			v31 = v30
		} else {
			v31 = v16
		}
		v32 = F_table_open(m, v25, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v34+v9<<(uint(int32(2))%32)))) = v32
			v39 = v32
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_InitResultRelInfo(m, l1, v39, l2, int32(0), v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				if v45 == int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v51 = F_palloc0(m, v48<<(uint(int32(2))%32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
						v54 = v51
						*(*int32)(unsafe.Add(mBase, uint32(v54+l2<<(uint(int32(2))%32)-int32(4)))) = l1
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v62 = F_lappend(m, v61, l1)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v62
							return
						}
					}
				} else {
					v54 = v45
					*(*int32)(unsafe.Add(mBase, uint32(v54+l2<<(uint(int32(2))%32)-int32(4)))) = l1
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v62 = F_lappend(m, v61, l1)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v62
						return
					}
				}
			}
		}
	} else {
		v39 = v13
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		F_InitResultRelInfo(m, l1, v39, l2, int32(0), v42)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v45 == int32(0) {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v51 = F_palloc0(m, v48<<(uint(int32(2))%32))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v51
					v54 = v51
					*(*int32)(unsafe.Add(mBase, uint32(v54+l2<<(uint(int32(2))%32)-int32(4)))) = l1
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v62 = F_lappend(m, v61, l1)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v62
						return
					}
				}
			} else {
				v54 = v45
				*(*int32)(unsafe.Add(mBase, uint32(v54+l2<<(uint(int32(2))%32)-int32(4)))) = l1
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v62 = F_lappend(m, v61, l1)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v62
					return
				}
			}
		}
	}
}
func F_ExecResult(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v12 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
			if v18 != int32(1) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				F_MemoryContextReset(m, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
					if v47 != 0 {
						v93 = v2
						m.G0 = v9 + int32(16)
						return v93
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v48 != 0 {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
							if v49 != 0 {
								F_ExecReScan(m, v48)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
									v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
											if v57&int32(2) != 0 {
												v93 = v2
												m.G0 = v9 + int32(16)
												return v93
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
												v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
												m.T0[v68].(func(*base.Module, int32))(m, v66)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = int32(4455216)
													v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
													v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
														v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
														v86 = v84 & int32(65533)
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
														v93 = v66
														m.G0 = v9 + int32(16)
														return v93
													}
												}
											}
										}
									}
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
								v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53 == int32(0) {
										v93 = v2
										m.G0 = v9 + int32(16)
										return v93
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
										if v57&int32(2) != 0 {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
											m.T0[v68].(func(*base.Module, int32))(m, v66)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = int32(4455216)
												v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
												v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
													v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
													v86 = v84 & int32(65533)
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
													v93 = v66
													m.G0 = v9 + int32(16)
													return v93
												}
											}
										}
									}
								}
							}
						} else {
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
							m.T0[v68].(func(*base.Module, int32))(m, v66)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(4455216)
								v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
								v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
									v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
									v86 = v84 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
									v93 = v66
									m.G0 = v9 + int32(16)
									return v93
								}
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v21 == int32(0) {
					v24 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v24)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
					F_MemoryContextReset(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
						if v47 != 0 {
							v93 = v2
							m.G0 = v9 + int32(16)
							return v93
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v48 != 0 {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
								if v49 != 0 {
									F_ExecReScan(m, v48)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
										v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											if v53 == int32(0) {
												v93 = v2
												m.G0 = v9 + int32(16)
												return v93
											} else {
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
												if v57&int32(2) != 0 {
													v93 = v2
													m.G0 = v9 + int32(16)
													return v93
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
													v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
													m.T0[v68].(func(*base.Module, int32))(m, v66)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = int32(4455216)
														v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
														v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
															v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
															v86 = v84 & int32(65533)
															*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
															*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
															v93 = v66
															m.G0 = v9 + int32(16)
															return v93
														}
													}
												}
											}
										}
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
									v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
											if v57&int32(2) != 0 {
												v93 = v2
												m.G0 = v9 + int32(16)
												return v93
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
												v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
												m.T0[v68].(func(*base.Module, int32))(m, v66)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = int32(4455216)
													v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
													v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
														v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
														v86 = v84 & int32(65533)
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
														v93 = v66
														m.G0 = v9 + int32(16)
														return v93
													}
												}
											}
										}
									}
								}
							} else {
								v61 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
								m.T0[v68].(func(*base.Module, int32))(m, v66)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = int32(4455216)
									v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
									v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
										v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
										v86 = v84 & int32(65533)
										*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
										*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
										v93 = v66
										m.G0 = v9 + int32(16)
										return v93
									}
								}
							}
						}
					}
				} else {
					v26 = int32(4455216)
					v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
					v34 = m.T0[v33].(func(*base.Module, int32, int32, int32) int32)(m, v21, v17, v9+int32(15))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v38)
						if v34 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
							F_MemoryContextReset(m, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
								if v47 != 0 {
									v93 = v2
									m.G0 = v9 + int32(16)
									return v93
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v48 != 0 {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
										if v49 != 0 {
											F_ExecReScan(m, v48)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
												v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													if v53 == int32(0) {
														v93 = v2
														m.G0 = v9 + int32(16)
														return v93
													} else {
														v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
														if v57&int32(2) != 0 {
															v93 = v2
															m.G0 = v9 + int32(16)
															return v93
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
															v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
															v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
															m.T0[v68].(func(*base.Module, int32))(m, v66)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v71 = int32(4455216)
																v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
																*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
																v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
																v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
																	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
																	v86 = v84 & int32(65533)
																	*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
																	v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
																	v93 = v66
																	m.G0 = v9 + int32(16)
																	return v93
																}
															}
														}
													}
												}
											}
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
											v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												if v53 == int32(0) {
													v93 = v2
													m.G0 = v9 + int32(16)
													return v93
												} else {
													v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
													if v57&int32(2) != 0 {
														v93 = v2
														m.G0 = v9 + int32(16)
														return v93
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
														v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
														v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
														m.T0[v68].(func(*base.Module, int32))(m, v66)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = int32(4455216)
															v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
															v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
																v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
																v86 = v84 & int32(65533)
																*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
																v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
																*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
																v93 = v66
																m.G0 = v9 + int32(16)
																return v93
															}
														}
													}
												}
											}
										}
									} else {
										v61 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
										m.T0[v68].(func(*base.Module, int32))(m, v66)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = int32(4455216)
											v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
											v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
												v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
												v86 = v84 & int32(65533)
												*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
												*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
												v93 = v66
												m.G0 = v9 + int32(16)
												return v93
											}
										}
									}
								}
							}
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v40)
							v93 = v2
							m.G0 = v9 + int32(16)
							return v93
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
		if v18 != int32(1) {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
			F_MemoryContextReset(m, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
				if v47 != 0 {
					v93 = v2
					m.G0 = v9 + int32(16)
					return v93
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v48 != 0 {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
						if v49 != 0 {
							F_ExecReScan(m, v48)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
								v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53 == int32(0) {
										v93 = v2
										m.G0 = v9 + int32(16)
										return v93
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
										if v57&int32(2) != 0 {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
											m.T0[v68].(func(*base.Module, int32))(m, v66)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = int32(4455216)
												v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
												v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
													v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
													v86 = v84 & int32(65533)
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
													v93 = v66
													m.G0 = v9 + int32(16)
													return v93
												}
											}
										}
									}
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
							v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v53 == int32(0) {
									v93 = v2
									m.G0 = v9 + int32(16)
									return v93
								} else {
									v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
									if v57&int32(2) != 0 {
										v93 = v2
										m.G0 = v9 + int32(16)
										return v93
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
										m.T0[v68].(func(*base.Module, int32))(m, v66)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = int32(4455216)
											v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
											v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
												v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
												v86 = v84 & int32(65533)
												*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
												*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
												v93 = v66
												m.G0 = v9 + int32(16)
												return v93
											}
										}
									}
								}
							}
						}
					} else {
						v61 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
						m.T0[v68].(func(*base.Module, int32))(m, v66)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = int32(4455216)
							v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
							v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
								v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
								v86 = v84 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
								*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
								v93 = v66
								m.G0 = v9 + int32(16)
								return v93
							}
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v21 == int32(0) {
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v24)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				F_MemoryContextReset(m, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
					if v47 != 0 {
						v93 = v2
						m.G0 = v9 + int32(16)
						return v93
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v48 != 0 {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
							if v49 != 0 {
								F_ExecReScan(m, v48)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
									v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
											if v57&int32(2) != 0 {
												v93 = v2
												m.G0 = v9 + int32(16)
												return v93
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
												v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
												m.T0[v68].(func(*base.Module, int32))(m, v66)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = int32(4455216)
													v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
													v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
														v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
														v86 = v84 & int32(65533)
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
														*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
														v93 = v66
														m.G0 = v9 + int32(16)
														return v93
													}
												}
											}
										}
									}
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
								v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53 == int32(0) {
										v93 = v2
										m.G0 = v9 + int32(16)
										return v93
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
										if v57&int32(2) != 0 {
											v93 = v2
											m.G0 = v9 + int32(16)
											return v93
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
											m.T0[v68].(func(*base.Module, int32))(m, v66)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = int32(4455216)
												v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
												*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
												v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
													v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
													v86 = v84 & int32(65533)
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
													*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
													v93 = v66
													m.G0 = v9 + int32(16)
													return v93
												}
											}
										}
									}
								}
							}
						} else {
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
							m.T0[v68].(func(*base.Module, int32))(m, v66)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(4455216)
								v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
								v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
									v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
									v86 = v84 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
									*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
									v93 = v66
									m.G0 = v9 + int32(16)
									return v93
								}
							}
						}
					}
				}
			} else {
				v26 = int32(4455216)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
				v34 = m.T0[v33].(func(*base.Module, int32, int32, int32) int32)(m, v21, v17, v9+int32(15))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v38)
					if v34 != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
						F_MemoryContextReset(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
							if v47 != 0 {
								v93 = v2
								m.G0 = v9 + int32(16)
								return v93
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v48 != 0 {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
									if v49 != 0 {
										F_ExecReScan(m, v48)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
											v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												if v53 == int32(0) {
													v93 = v2
													m.G0 = v9 + int32(16)
													return v93
												} else {
													v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
													if v57&int32(2) != 0 {
														v93 = v2
														m.G0 = v9 + int32(16)
														return v93
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
														v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
														v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
														m.T0[v68].(func(*base.Module, int32))(m, v66)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = int32(4455216)
															v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
															v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
																v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
																v86 = v84 & int32(65533)
																*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
																v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
																*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
																v93 = v66
																m.G0 = v9 + int32(16)
																return v93
															}
														}
													}
												}
											}
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
										v53 = m.T0[v52].(func(*base.Module, int32) int32)(m, v48)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											if v53 == int32(0) {
												v93 = v2
												m.G0 = v9 + int32(16)
												return v93
											} else {
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
												if v57&int32(2) != 0 {
													v93 = v2
													m.G0 = v9 + int32(16)
													return v93
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v53
													v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
													m.T0[v68].(func(*base.Module, int32))(m, v66)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = int32(4455216)
														v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
														v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
															v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
															v86 = v84 & int32(65533)
															*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
															*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
															v93 = v66
															m.G0 = v9 + int32(16)
															return v93
														}
													}
												}
											}
										}
									}
								} else {
									v61 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v61)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
									m.T0[v68].(func(*base.Module, int32))(m, v66)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = int32(4455216)
										v72 = *(*int32)(unsafe.Add(mBase, _consts[0]))
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
										v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v64+int32(4), v65, int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v72
											v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
											v86 = v84 & int32(65533)
											*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v86)
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
											*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)) = uint16(v89)
											v93 = v66
											m.G0 = v9 + int32(16)
											return v93
										}
									}
								}
							}
						}
					} else {
						v40 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v40)
						v93 = v2
						m.G0 = v9 + int32(16)
						return v93
					}
				}
			}
		}
	}
}
