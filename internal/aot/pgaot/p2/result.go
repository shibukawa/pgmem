package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitResultRelation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = l2 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(int32(2))%32))))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+l2<<(uint(int32(2))%32)-int32(4))))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitResultRelation[0]))
		if int32(0) <= v25 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
			v30 = v28
		} else {
			v30 = int32(0)
		}
		v31 = F_table_open(m, v23, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v33+v8<<(uint(int32(2))%32)))) = v31
			v38 = v31
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			F_InitResultRelInfo(m, l1, v38, l2, int32(0), v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				if v43 == int32(0) {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v49 = F_palloc0(m, v46<<(uint(int32(2))%32))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v49
						v52 = v49
						*(*int32)(unsafe.Add(mBase, uint32(v52+l2<<(uint(int32(2))%32)-int32(4)))) = l1
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v60 = F_lappend(m, v59, l1)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v60
							return
						}
					}
				} else {
					v52 = v43
					*(*int32)(unsafe.Add(mBase, uint32(v52+l2<<(uint(int32(2))%32)-int32(4)))) = l1
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v60 = F_lappend(m, v59, l1)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v60
						return
					}
				}
			}
		}
	} else {
		v38 = v12
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		F_InitResultRelInfo(m, l1, v38, l2, int32(0), v40)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v43 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v49 = F_palloc0(m, v46<<(uint(int32(2))%32))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v49
					v52 = v49
					*(*int32)(unsafe.Add(mBase, uint32(v52+l2<<(uint(int32(2))%32)-int32(4)))) = l1
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v60 = F_lappend(m, v59, l1)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v60
						return
					}
				}
			} else {
				v52 = v43
				*(*int32)(unsafe.Add(mBase, uint32(v52+l2<<(uint(int32(2))%32)-int32(4)))) = l1
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v60 = F_lappend(m, v59, l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v60
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[0]))
	if v13 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
			if v19 != int32(1) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				F_MemoryContextReset(m, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
					if v50 != 0 {
						v97 = v2
						m.G0 = v10 + int32(16)
						return v97
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v51 != 0 {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
							if v52 != 0 {
								F_ExecReScan(m, v51)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
									v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 == int32(0) {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
											if v60&int32(2) != 0 {
												v97 = v2
												m.G0 = v10 + int32(16)
												return v97
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
												m.T0[v71].(func(*base.Module, int32))(m, v69)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = int32(_a_F_ExecResult_0)
													v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
													v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
														v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
														v89 = v87 & int32(_a_F_ExecResult_1)
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
														v97 = v69
														m.G0 = v10 + int32(16)
														return v97
													}
												}
											}
										}
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 == int32(0) {
										v97 = v2
										m.G0 = v10 + int32(16)
										return v97
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
										if v60&int32(2) != 0 {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
											m.T0[v71].(func(*base.Module, int32))(m, v69)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = int32(_a_F_ExecResult_0)
												v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
												*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
												v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
													v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
													v89 = v87 & int32(_a_F_ExecResult_1)
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
													v97 = v69
													m.G0 = v10 + int32(16)
													return v97
												}
											}
										}
									}
								}
							}
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
							m.T0[v71].(func(*base.Module, int32))(m, v69)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = int32(_a_F_ExecResult_0)
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
								v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
									v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
									v89 = v87 & int32(_a_F_ExecResult_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
									v97 = v69
									m.G0 = v10 + int32(16)
									return v97
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v22 == int32(0) {
					v25 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v25)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
					F_MemoryContextReset(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
						if v50 != 0 {
							v97 = v2
							m.G0 = v10 + int32(16)
							return v97
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v51 != 0 {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
								if v52 != 0 {
									F_ExecReScan(m, v51)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
										v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 == int32(0) {
												v97 = v2
												m.G0 = v10 + int32(16)
												return v97
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
												if v60&int32(2) != 0 {
													v97 = v2
													m.G0 = v10 + int32(16)
													return v97
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
													v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
													m.T0[v71].(func(*base.Module, int32))(m, v69)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = int32(_a_F_ExecResult_0)
														v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
														*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
														v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
															v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
															v89 = v87 & int32(_a_F_ExecResult_1)
															*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
															*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
															v97 = v69
															m.G0 = v10 + int32(16)
															return v97
														}
													}
												}
											}
										}
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
									v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 == int32(0) {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
											if v60&int32(2) != 0 {
												v97 = v2
												m.G0 = v10 + int32(16)
												return v97
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
												m.T0[v71].(func(*base.Module, int32))(m, v69)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = int32(_a_F_ExecResult_0)
													v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
													v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
														v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
														v89 = v87 & int32(_a_F_ExecResult_1)
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
														v97 = v69
														m.G0 = v10 + int32(16)
														return v97
													}
												}
											}
										}
									}
								}
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
								m.T0[v71].(func(*base.Module, int32))(m, v69)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = int32(_a_F_ExecResult_0)
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
									v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
										v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
										v89 = v87 & int32(_a_F_ExecResult_1)
										*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
										*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
										v97 = v69
										m.G0 = v10 + int32(16)
										return v97
									}
								}
							}
						}
					}
				} else {
					v27 = int32(_a_F_ExecResult_0)
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v30
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					v35 = m.T0[v34].(func(*base.Module, int32, int32, int32) int64)(m, v22, v18, v10+int32(15))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v28
						v39 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v39)
						if v35 != int64(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							F_MemoryContextReset(m, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
								if v50 != 0 {
									v97 = v2
									m.G0 = v10 + int32(16)
									return v97
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v51 != 0 {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
										if v52 != 0 {
											F_ExecReScan(m, v51)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
												v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													if v56 == int32(0) {
														v97 = v2
														m.G0 = v10 + int32(16)
														return v97
													} else {
														v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
														if v60&int32(2) != 0 {
															v97 = v2
															m.G0 = v10 + int32(16)
															return v97
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
															v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
															m.T0[v71].(func(*base.Module, int32))(m, v69)
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																v74 = int32(_a_F_ExecResult_0)
																v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
																v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
																*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
																v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
																	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
																	v89 = v87 & int32(_a_F_ExecResult_1)
																	*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
																	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
																	v97 = v69
																	m.G0 = v10 + int32(16)
																	return v97
																}
															}
														}
													}
												}
											}
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
											v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 == int32(0) {
													v97 = v2
													m.G0 = v10 + int32(16)
													return v97
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
													if v60&int32(2) != 0 {
														v97 = v2
														m.G0 = v10 + int32(16)
														return v97
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
														v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
														m.T0[v71].(func(*base.Module, int32))(m, v69)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															v74 = int32(_a_F_ExecResult_0)
															v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
															v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
															v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
																v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
																v89 = v87 & int32(_a_F_ExecResult_1)
																*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
																*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
																v97 = v69
																m.G0 = v10 + int32(16)
																return v97
															}
														}
													}
												}
											}
										}
									} else {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
										m.T0[v71].(func(*base.Module, int32))(m, v69)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = int32(_a_F_ExecResult_0)
											v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
											*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
											v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
												v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
												v89 = v87 & int32(_a_F_ExecResult_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
												*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
												v97 = v69
												m.G0 = v10 + int32(16)
												return v97
											}
										}
									}
								}
							}
						} else {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v43)
							v97 = v2
							m.G0 = v10 + int32(16)
							return v97
						}
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)))
		if v19 != int32(1) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
			F_MemoryContextReset(m, v47)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
				if v50 != 0 {
					v97 = v2
					m.G0 = v10 + int32(16)
					return v97
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v51 != 0 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
						if v52 != 0 {
							F_ExecReScan(m, v51)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 == int32(0) {
										v97 = v2
										m.G0 = v10 + int32(16)
										return v97
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
										if v60&int32(2) != 0 {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
											m.T0[v71].(func(*base.Module, int32))(m, v69)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = int32(_a_F_ExecResult_0)
												v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
												*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
												v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
													v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
													v89 = v87 & int32(_a_F_ExecResult_1)
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
													v97 = v69
													m.G0 = v10 + int32(16)
													return v97
												}
											}
										}
									}
								}
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
							v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 == int32(0) {
									v97 = v2
									m.G0 = v10 + int32(16)
									return v97
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
									if v60&int32(2) != 0 {
										v97 = v2
										m.G0 = v10 + int32(16)
										return v97
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
										m.T0[v71].(func(*base.Module, int32))(m, v69)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = int32(_a_F_ExecResult_0)
											v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
											*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
											v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
												v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
												v89 = v87 & int32(_a_F_ExecResult_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
												*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
												v97 = v69
												m.G0 = v10 + int32(16)
												return v97
											}
										}
									}
								}
							}
						}
					} else {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
						m.T0[v71].(func(*base.Module, int32))(m, v69)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = int32(_a_F_ExecResult_0)
							v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
							v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
								v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
								v89 = v87 & int32(_a_F_ExecResult_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
								v97 = v69
								m.G0 = v10 + int32(16)
								return v97
							}
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v22 == int32(0) {
				v25 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v25)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				F_MemoryContextReset(m, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
					if v50 != 0 {
						v97 = v2
						m.G0 = v10 + int32(16)
						return v97
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v51 != 0 {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
							if v52 != 0 {
								F_ExecReScan(m, v51)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
									v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 == int32(0) {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
											if v60&int32(2) != 0 {
												v97 = v2
												m.G0 = v10 + int32(16)
												return v97
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
												m.T0[v71].(func(*base.Module, int32))(m, v69)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = int32(_a_F_ExecResult_0)
													v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
													v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
														v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
														v89 = v87 & int32(_a_F_ExecResult_1)
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
														*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
														v97 = v69
														m.G0 = v10 + int32(16)
														return v97
													}
												}
											}
										}
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 == int32(0) {
										v97 = v2
										m.G0 = v10 + int32(16)
										return v97
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
										if v60&int32(2) != 0 {
											v97 = v2
											m.G0 = v10 + int32(16)
											return v97
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
											m.T0[v71].(func(*base.Module, int32))(m, v69)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = int32(_a_F_ExecResult_0)
												v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
												*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
												v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
													v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
													v89 = v87 & int32(_a_F_ExecResult_1)
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
													*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
													v97 = v69
													m.G0 = v10 + int32(16)
													return v97
												}
											}
										}
									}
								}
							}
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
							m.T0[v71].(func(*base.Module, int32))(m, v69)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = int32(_a_F_ExecResult_0)
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
								v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
									v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
									v89 = v87 & int32(_a_F_ExecResult_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
									v97 = v69
									m.G0 = v10 + int32(16)
									return v97
								}
							}
						}
					}
				}
			} else {
				v27 = int32(_a_F_ExecResult_0)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v30
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
				v35 = m.T0[v34].(func(*base.Module, int32, int32, int32) int64)(m, v22, v18, v10+int32(15))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v28
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+109)) = uint8(v39)
					if v35 != int64(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
						F_MemoryContextReset(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
							if v50 != 0 {
								v97 = v2
								m.G0 = v10 + int32(16)
								return v97
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v51 != 0 {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
									if v52 != 0 {
										F_ExecReScan(m, v51)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
											v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 == int32(0) {
													v97 = v2
													m.G0 = v10 + int32(16)
													return v97
												} else {
													v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
													if v60&int32(2) != 0 {
														v97 = v2
														m.G0 = v10 + int32(16)
														return v97
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
														v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
														m.T0[v71].(func(*base.Module, int32))(m, v69)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															v74 = int32(_a_F_ExecResult_0)
															v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
															v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
															*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
															v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
																v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
																v89 = v87 & int32(_a_F_ExecResult_1)
																*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
																*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
																v97 = v69
																m.G0 = v10 + int32(16)
																return v97
															}
														}
													}
												}
											}
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
										v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v51)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 == int32(0) {
												v97 = v2
												m.G0 = v10 + int32(16)
												return v97
											} else {
												v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
												if v60&int32(2) != 0 {
													v97 = v2
													m.G0 = v10 + int32(16)
													return v97
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v56
													v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
													m.T0[v71].(func(*base.Module, int32))(m, v69)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = int32(_a_F_ExecResult_0)
														v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
														*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
														v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
															v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
															v89 = v87 & int32(_a_F_ExecResult_1)
															*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
															*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
															v97 = v69
															m.G0 = v10 + int32(16)
															return v97
														}
													}
												}
											}
										}
									}
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v64)
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
									m.T0[v71].(func(*base.Module, int32))(m, v69)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = int32(_a_F_ExecResult_0)
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1]))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
										*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v77
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
										v83 = m.T0[v82].(func(*base.Module, int32, int32, int32) int64)(m, v67+int32(8), v68, int32(0))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ExecResult[1])) = v75
											v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
											v89 = v87 & int32(_a_F_ExecResult_1)
											*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)) = uint16(v89)
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
											*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v92)
											v97 = v69
											m.G0 = v10 + int32(16)
											return v97
										}
									}
								}
							}
						}
					} else {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)) = uint8(v43)
						v97 = v2
						m.G0 = v10 + int32(16)
						return v97
					}
				}
			}
		}
	}
}
