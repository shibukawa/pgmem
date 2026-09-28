package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_or(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14393(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_tsquery_phrase_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_copy(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			if base.Ui32(base.I32_wrap_i64(v20)) < base.Ui32(int32(_a_F_tsquery_phrase_distance_0)) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v24 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v13 != v27 {
						v89 = v18
						v91 = v13
						F_pfree(m, v91)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							v94 = v89
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v94)
						}
					} else {
						v94 = v18
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v94)
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					if v29 == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v32 != v18 {
							v89 = v13
							v91 = v18
							F_pfree(m, v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v94 = v89
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v94)
							}
						} else {
							v94 = v13
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v94)
						}
					} else {
						v35 = F_palloc0(m, int32(24))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v37 | int32(1)
							v42 = F_palloc0(m, int32(12))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35))) = v42
								v45 = int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v45)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								v48 = int32(4)
								*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)) = uint8(v48)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)) = uint16(v20)
								v54 = F_palloc0_mul(m, v48, v45)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v54
									v58 = v18 + int32(8)
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
									v63 = F_QT2QTN(m, v58, v58+v59*int32(12))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int64(0)
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v65))) = v63
										v68 = v13 + int32(8)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v73 = F_QT2QTN(m, v68, v68+v69*int32(12))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int64(0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(2)
											v79 = F_QTN2QT(m, v35)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int64(0)
											} else {
												F_QTNFree(m, v35)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													if v83 != v13 {
														F_pfree(m, v13)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int64(0)
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															if v87 == v18 {
																v94 = v79
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v94)
															} else {
																v89 = v79
																v91 = v18
																F_pfree(m, v91)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	v94 = v89
																	m.G0 = v10 + int32(16)
																	return base.I64_extend_i32_u(v94)
																}
															}
														}
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														if v87 == v18 {
															v94 = v79
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v94)
														} else {
															v89 = v79
															v91 = v18
															F_pfree(m, v91)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																v94 = v89
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v94)
															}
														}
													}
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_tsquery_phrase_distance_1)
						F_errmsg(m, int32(_a_F_tsquery_phrase_distance_2), v10)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_tsquery_phrase_distance_3), int32(126), int32(_a_F_tsquery_phrase_distance_4))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
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
	}
}
