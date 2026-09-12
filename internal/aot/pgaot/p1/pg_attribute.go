package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_attribute_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 == int32(0) {
			if l4 != 0 {
				v20 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v20)
				v128 = int64(0)
				m.G0 = v11 + int32(48)
				return v128
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50360452))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
						F_errmsg(m, int32(70656), v11)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(509618), int32(3190), int32(65010))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41)+91)))
			if v43 == int32(1) {
				if l4 != 0 {
					v46 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v46)
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int64(0)
					} else {
						v128 = int64(0)
						m.G0 = v11 + int32(48)
						return v128
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50360452))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
							F_errmsg(m, int32(70656), v11+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(509618), int32(3209), int32(65010))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v71 = F_SysCacheGetAttr(m, int32(7), v14, int32(22), v11+int32(47))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int64(0)
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
					if v73 != int32(1) {
						v81 = F_SearchSysCache1(m, int32(57), l0)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							if v81 == int32(0) {
								F_ReleaseCatCache(m, v14)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
								} else {
									if l4 != 0 {
										v87 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v87)
										v128 = int64(0)
										m.G0 = v11 + int32(48)
										return v128
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(16908420))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
												F_errmsg(m, int32(70672), v11+int32(32))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(509618), int32(3247), int32(65010))
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								}
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
								v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+22)))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109)+80))
								F_ReleaseCatCache(m, v81)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									v114 = F_pg_detoast_datum(m, v71)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int64(0)
									} else {
										v117 = F_aclmask(m, v114, l2, v111, l3, int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int64(0)
										} else {
											if v114 == int32(0) {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int64(0)
												} else {
													v128 = v117
													m.G0 = v11 + int32(48)
													return v128
												}
											} else {
												if v71 == v114 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int64(0)
													} else {
														v128 = v117
														m.G0 = v11 + int32(48)
														return v128
													}
												} else {
													F_pfree(m, v114)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int64(0)
														} else {
															v128 = v117
															m.G0 = v11 + int32(48)
															return v128
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_ReleaseCatCache(m, v14)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							v128 = int64(0)
							m.G0 = v11 + int32(48)
							return v128
						}
					}
				}
			}
		}
	}
}
