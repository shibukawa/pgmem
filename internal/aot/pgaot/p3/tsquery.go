package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_or(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14010(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_tsquery_phrase_distance(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_copy(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if base.Ui32(v20) < base.Ui32(int32(_a_F_tsquery_phrase_distance_0)) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v23 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v13 != v26 {
						v87 = v18
						v89 = v13
						F_pfree(m, v89)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = v87
							m.G0 = v10 + int32(16)
							return v93
						}
					} else {
						v93 = v18
						m.G0 = v10 + int32(16)
						return v93
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					if v28 == int32(0) {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v31 != v18 {
							v87 = v13
							v89 = v18
							F_pfree(m, v89)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = v87
								m.G0 = v10 + int32(16)
								return v93
							}
						} else {
							v93 = v13
							m.G0 = v10 + int32(16)
							return v93
						}
					} else {
						v34 = F_palloc0(m, int32(24))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v36 | int32(1)
							v41 = F_palloc0(m, int32(12))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = v41
								v44 = int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v44)
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
								v47 = int32(4)
								*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)) = uint8(v47)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
								*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)) = uint16(v20)
								v52 = F_palloc0(m, int32(8))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v52
									v56 = v18 + int32(8)
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
									v61 = F_QT2QTN(m, v56, v56+v57*int32(12))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v63))) = v61
										v66 = v13 + int32(8)
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v71 = F_QT2QTN(m, v66, v66+v67*int32(12))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(2)
											v77 = F_QTN2QT(m, v34)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_QTNFree(m, v34)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													if v81 != v13 {
														F_pfree(m, v13)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															if v85 == v18 {
																v93 = v77
																m.G0 = v10 + int32(16)
																return v93
															} else {
																v87 = v77
																v89 = v18
																F_pfree(m, v89)
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	v93 = v87
																	m.G0 = v10 + int32(16)
																	return v93
																}
															}
														}
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v85 == v18 {
															v93 = v77
															m.G0 = v10 + int32(16)
															return v93
														} else {
															v87 = v77
															v89 = v18
															F_pfree(m, v89)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																v93 = v87
																m.G0 = v10 + int32(16)
																return v93
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
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_tsquery_phrase_distance_1)
						F_errmsg(m, int32(_a_F_tsquery_phrase_distance_2), v10)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_tsquery_phrase_distance_3), int32(126), int32(_a_F_tsquery_phrase_distance_4))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
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
