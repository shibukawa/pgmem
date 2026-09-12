package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_copy(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_copy(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v19 != v7 {
					v78 = v7
					v79 = v14
					F_pfree(m, v78)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v83 = v79
						return v83
					}
				} else {
					v83 = v14
					return v83
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v21 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v14 != v24 {
						v78 = v14
						v79 = v7
						F_pfree(m, v78)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v83 = v79
							return v83
						}
					} else {
						v83 = v7
						return v83
					}
				} else {
					v27 = F_palloc0(m, int32(24))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v29 | int32(1)
						v34 = F_palloc0(m, int32(12))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27))) = v34
							v37 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v37)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							v40 = int32(3)
							*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)) = uint8(v40)
							v43 = F_palloc0(m, int32(8))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v43
								v47 = v14 + int32(8)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
								v52 = F_QT2QTN(m, v47, v47+v48*int32(12))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v54))) = v52
									v57 = v7 + int32(8)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
									v62 = F_QT2QTN(m, v57, v57+v58*int32(12))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v62
										*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(2)
										v68 = F_QTN2QT(m, v27)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											F_QTNFree(m, v27)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v72 != v7 {
													F_pfree(m, v7)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v14 == v76 {
															v83 = v68
															return v83
														} else {
															v78 = v14
															v79 = v68
															F_pfree(m, v78)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																v83 = v79
																return v83
															}
														}
													}
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v14 == v76 {
														v83 = v68
														return v83
													} else {
														v78 = v14
														v79 = v68
														F_pfree(m, v78)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v83 = v79
															return v83
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
		}
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
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
			if base.Ui32(v20) < base.Ui32(int32(16385)) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v23 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v13 != v26 {
						v87 = v13
						v88 = v18
						F_pfree(m, v87)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v93 = v88
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
						if v18 != v31 {
							v87 = v18
							v88 = v13
							F_pfree(m, v87)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v93 = v88
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
															if v18 == v85 {
																v93 = v77
																m.G0 = v10 + int32(16)
																return v93
															} else {
																v87 = v18
																v88 = v77
																F_pfree(m, v87)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	v93 = v88
																	m.G0 = v10 + int32(16)
																	return v93
																}
															}
														}
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v18 == v85 {
															v93 = v77
															m.G0 = v10 + int32(16)
															return v93
														} else {
															v87 = v18
															v88 = v77
															F_pfree(m, v87)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																v93 = v88
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
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16384)
						F_errmsg(m, int32(326627), v10)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472616), int32(126), int32(397536))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
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
