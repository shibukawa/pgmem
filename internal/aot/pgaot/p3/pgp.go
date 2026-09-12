package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_disable_mdc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = base.B2i32(l1 != v3)
	return v3
}
func F_pgp_elgamal_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = int32(-109)
	v16 = F_mpi_check(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v92 = v15
			return v92
		} else {
			v22 = F_mpi_check(m, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v92 = v15
					return v92
				} else {
					v26 = F_mpi_check(m, v13)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 == int32(0) {
							v92 = v15
							return v92
						} else {
							v30 = F_mpi_check(m, v12)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								if v30 == int32(0) {
									v92 = v15
									return v92
								} else {
									v34 = int32(1)
									if v14 <= v34 {
										v37 = v34
									} else {
										v37 = v14
									}
									v38 = F_palloc(m, v37)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v40 = F_palloc(m, v37)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = F_palloc(m, v37)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int32(0)
											} else {
												v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
												v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												v51 = m.Env.Pgmem_bn_op(m, int32(1), v45, v46, v47, v48, v49, v50, v38, v14)
												mBase = m.M
												if v51 < int32(0) {
													v76 = v15
													if v42 != 0 {
														v79 = F___memset(m, v42, int32(0), v37)
														mBase = m.M
														F_pfree(m, v42)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															if v40 != 0 {
																v83 = F___memset(m, v40, int32(0), v37)
																mBase = m.M
																F_pfree(m, v40)
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return int32(0)
																} else {
																	if v38 == int32(0) {
																		v92 = v76
																		return v92
																	} else {
																		v89 = F___memset(m, v38, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v38)
																		mBase = m.M
																		v91 = m.ExcPending
																		if v91 != 0 {
																			return int32(0)
																		} else {
																			v92 = v76
																			return v92
																		}
																	}
																}
															} else {
																if v38 == int32(0) {
																	v92 = v76
																	return v92
																} else {
																	v89 = F___memset(m, v38, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v38)
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int32(0)
																	} else {
																		v92 = v76
																		return v92
																	}
																}
															}
														}
													} else {
														if v40 != 0 {
															v83 = F___memset(m, v40, int32(0), v37)
															mBase = m.M
															F_pfree(m, v40)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																if v38 == int32(0) {
																	v92 = v76
																	return v92
																} else {
																	v89 = F___memset(m, v38, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v38)
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int32(0)
																	} else {
																		v92 = v76
																		return v92
																	}
																}
															}
														} else {
															if v38 == int32(0) {
																v92 = v76
																return v92
															} else {
																v89 = F___memset(m, v38, int32(0), v37)
																mBase = m.M
																F_pfree(m, v38)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	v92 = v76
																	return v92
																}
															}
														}
													}
												} else {
													v55 = int32(0)
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													v59 = m.Env.Pgmem_bn_op(m, int32(3), v38, v51, v55, v55, v57, v58, v40, v14)
													mBase = m.M
													if v59 < v55 {
														v76 = v15
														if v42 != 0 {
															v79 = F___memset(m, v42, int32(0), v37)
															mBase = m.M
															F_pfree(m, v42)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																if v40 != 0 {
																	v83 = F___memset(m, v40, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v40)
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return int32(0)
																	} else {
																		if v38 == int32(0) {
																			v92 = v76
																			return v92
																		} else {
																			v89 = F___memset(m, v38, int32(0), v37)
																			mBase = m.M
																			F_pfree(m, v38)
																			mBase = m.M
																			v91 = m.ExcPending
																			if v91 != 0 {
																				return int32(0)
																			} else {
																				v92 = v76
																				return v92
																			}
																		}
																	}
																} else {
																	if v38 == int32(0) {
																		v92 = v76
																		return v92
																	} else {
																		v89 = F___memset(m, v38, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v38)
																		mBase = m.M
																		v91 = m.ExcPending
																		if v91 != 0 {
																			return int32(0)
																		} else {
																			v92 = v76
																			return v92
																		}
																	}
																}
															}
														} else {
															if v40 != 0 {
																v83 = F___memset(m, v40, int32(0), v37)
																mBase = m.M
																F_pfree(m, v40)
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return int32(0)
																} else {
																	if v38 == int32(0) {
																		v92 = v76
																		return v92
																	} else {
																		v89 = F___memset(m, v38, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v38)
																		mBase = m.M
																		v91 = m.ExcPending
																		if v91 != 0 {
																			return int32(0)
																		} else {
																			v92 = v76
																			return v92
																		}
																	}
																}
															} else {
																if v38 == int32(0) {
																	v92 = v76
																	return v92
																} else {
																	v89 = F___memset(m, v38, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v38)
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int32(0)
																	} else {
																		v92 = v76
																		return v92
																	}
																}
															}
														}
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
														v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														v67 = m.Env.Pgmem_bn_op(m, int32(2), v63, v64, v40, v59, v65, v66, v42, v14)
														mBase = m.M
														if v67 < int32(0) {
															v76 = v15
															if v42 != 0 {
																v79 = F___memset(m, v42, int32(0), v37)
																mBase = m.M
																F_pfree(m, v42)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	if v40 != 0 {
																		v83 = F___memset(m, v40, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v40)
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return int32(0)
																		} else {
																			if v38 == int32(0) {
																				v92 = v76
																				return v92
																			} else {
																				v89 = F___memset(m, v38, int32(0), v37)
																				mBase = m.M
																				F_pfree(m, v38)
																				mBase = m.M
																				v91 = m.ExcPending
																				if v91 != 0 {
																					return int32(0)
																				} else {
																					v92 = v76
																					return v92
																				}
																			}
																		}
																	} else {
																		if v38 == int32(0) {
																			v92 = v76
																			return v92
																		} else {
																			v89 = F___memset(m, v38, int32(0), v37)
																			mBase = m.M
																			F_pfree(m, v38)
																			mBase = m.M
																			v91 = m.ExcPending
																			if v91 != 0 {
																				return int32(0)
																			} else {
																				v92 = v76
																				return v92
																			}
																		}
																	}
																}
															} else {
																if v40 != 0 {
																	v83 = F___memset(m, v40, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v40)
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return int32(0)
																	} else {
																		if v38 == int32(0) {
																			v92 = v76
																			return v92
																		} else {
																			v89 = F___memset(m, v38, int32(0), v37)
																			mBase = m.M
																			F_pfree(m, v38)
																			mBase = m.M
																			v91 = m.ExcPending
																			if v91 != 0 {
																				return int32(0)
																			} else {
																				v92 = v76
																				return v92
																			}
																		}
																	}
																} else {
																	if v38 == int32(0) {
																		v92 = v76
																		return v92
																	} else {
																		v89 = F___memset(m, v38, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v38)
																		mBase = m.M
																		v91 = m.ExcPending
																		if v91 != 0 {
																			return int32(0)
																		} else {
																			v92 = v76
																			return v92
																		}
																	}
																}
															}
														} else {
															v70 = F_bytes_to_mpi(m, v42, v67)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
																if v70 != 0 {
																	v75 = int32(0)
																} else {
																	v75 = int32(-109)
																}
																v76 = v75
																if v42 != 0 {
																	v79 = F___memset(m, v42, int32(0), v37)
																	mBase = m.M
																	F_pfree(m, v42)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		if v40 != 0 {
																			v83 = F___memset(m, v40, int32(0), v37)
																			mBase = m.M
																			F_pfree(m, v40)
																			mBase = m.M
																			v85 = m.ExcPending
																			if v85 != 0 {
																				return int32(0)
																			} else {
																				if v38 == int32(0) {
																					v92 = v76
																					return v92
																				} else {
																					v89 = F___memset(m, v38, int32(0), v37)
																					mBase = m.M
																					F_pfree(m, v38)
																					mBase = m.M
																					v91 = m.ExcPending
																					if v91 != 0 {
																						return int32(0)
																					} else {
																						v92 = v76
																						return v92
																					}
																				}
																			}
																		} else {
																			if v38 == int32(0) {
																				v92 = v76
																				return v92
																			} else {
																				v89 = F___memset(m, v38, int32(0), v37)
																				mBase = m.M
																				F_pfree(m, v38)
																				mBase = m.M
																				v91 = m.ExcPending
																				if v91 != 0 {
																					return int32(0)
																				} else {
																					v92 = v76
																					return v92
																				}
																			}
																		}
																	}
																} else {
																	if v40 != 0 {
																		v83 = F___memset(m, v40, int32(0), v37)
																		mBase = m.M
																		F_pfree(m, v40)
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return int32(0)
																		} else {
																			if v38 == int32(0) {
																				v92 = v76
																				return v92
																			} else {
																				v89 = F___memset(m, v38, int32(0), v37)
																				mBase = m.M
																				F_pfree(m, v38)
																				mBase = m.M
																				v91 = m.ExcPending
																				if v91 != 0 {
																					return int32(0)
																				} else {
																					v92 = v76
																					return v92
																				}
																			}
																		}
																	} else {
																		if v38 == int32(0) {
																			v92 = v76
																			return v92
																		} else {
																			v89 = F___memset(m, v38, int32(0), v37)
																			mBase = m.M
																			F_pfree(m, v38)
																			mBase = m.M
																			v91 = m.ExcPending
																			if v91 != 0 {
																				return int32(0)
																			} else {
																				v92 = v76
																				return v92
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
				}
			}
		}
	}
}
func F_pgp_elgamal_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v20 = int32(-109)
	v21 = F_mpi_check(m, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		if v21 == int32(0) {
			v147 = v20
			return v147
		} else {
			v27 = F_mpi_check(m, v18)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v147 = v20
					return v147
				} else {
					v31 = F_mpi_check(m, v17)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							v147 = v20
							return v147
						} else {
							v35 = F_mpi_check(m, v16)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									v147 = v20
									return v147
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
									if v39 <= int32(5120) {
										v43 = base.I32_div_s(v39, int32(10))
										v54 = v43 + int32(160)
									} else {
										v46 = int32(3)
										v54 = int32(base.Ui32(int32(base.Ui32(v39)>>(uint(v46)%32))*v46+int32(600)) >> (uint(int32(1)) % 32))
									}
									v55 = int32(1)
									v59 = base.I32_div_s(v54+int32(7), int32(8))
									if v59 <= v55 {
										v62 = v55
									} else {
										v62 = v59
									}
									v63 = F_palloc(m, v62)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v65 = m.Env.Pgmem_bn_rand(m, v54, v63, v59)
										mBase = m.M
										if v65 < int32(0) {
											v133 = int32(-109)
											if v63 == int32(0) {
												v147 = v133
												return v147
											} else {
												v144 = F___memset(m, v63, int32(0), v62)
												mBase = m.M
												F_pfree(m, v63)
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int32(0)
												} else {
													v147 = v133
													return v147
												}
											}
										} else {
											v69 = int32(1)
											if v19 <= v69 {
												v72 = v69
											} else {
												v72 = v19
											}
											v73 = F_palloc(m, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = F_palloc(m, v72)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = F_palloc(m, v72)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														v79 = int32(-109)
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
														v85 = m.Env.Pgmem_bn_op(m, int32(1), v81, v82, v63, v65, v83, v84, v73, v19)
														mBase = m.M
														if v85 < int32(0) {
															v116 = v79
															if v77 != 0 {
																v120 = F___memset(m, v77, int32(0), v72)
																mBase = m.M
																F_pfree(m, v77)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int32(0)
																} else {
																	if v75 != 0 {
																		v124 = F___memset(m, v75, int32(0), v72)
																		mBase = m.M
																		F_pfree(m, v75)
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int32(0)
																		} else {
																			if v73 == int32(0) {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			} else {
																				v130 = F___memset(m, v73, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v73)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return int32(0)
																				} else {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				}
																			}
																		}
																	} else {
																		if v73 == int32(0) {
																			v133 = v116
																			if v63 == int32(0) {
																				v147 = v133
																				return v147
																			} else {
																				v144 = F___memset(m, v63, int32(0), v62)
																				mBase = m.M
																				F_pfree(m, v63)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					v147 = v133
																					return v147
																				}
																			}
																		} else {
																			v130 = F___memset(m, v73, int32(0), v72)
																			mBase = m.M
																			F_pfree(m, v73)
																			mBase = m.M
																			v132 = m.ExcPending
																			if v132 != 0 {
																				return int32(0)
																			} else {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																if v75 != 0 {
																	v124 = F___memset(m, v75, int32(0), v72)
																	mBase = m.M
																	F_pfree(m, v75)
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return int32(0)
																	} else {
																		if v73 == int32(0) {
																			v133 = v116
																			if v63 == int32(0) {
																				v147 = v133
																				return v147
																			} else {
																				v144 = F___memset(m, v63, int32(0), v62)
																				mBase = m.M
																				F_pfree(m, v63)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					v147 = v133
																					return v147
																				}
																			}
																		} else {
																			v130 = F___memset(m, v73, int32(0), v72)
																			mBase = m.M
																			F_pfree(m, v73)
																			mBase = m.M
																			v132 = m.ExcPending
																			if v132 != 0 {
																				return int32(0)
																			} else {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			}
																		}
																	}
																} else {
																	if v73 == int32(0) {
																		v133 = v116
																		if v63 == int32(0) {
																			v147 = v133
																			return v147
																		} else {
																			v144 = F___memset(m, v63, int32(0), v62)
																			mBase = m.M
																			F_pfree(m, v63)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				v147 = v133
																				return v147
																			}
																		}
																	} else {
																		v130 = F___memset(m, v73, int32(0), v72)
																		mBase = m.M
																		F_pfree(m, v73)
																		mBase = m.M
																		v132 = m.ExcPending
																		if v132 != 0 {
																			return int32(0)
																		} else {
																			v133 = v116
																			if v63 == int32(0) {
																				v147 = v133
																				return v147
																			} else {
																				v144 = F___memset(m, v63, int32(0), v62)
																				mBase = m.M
																				F_pfree(m, v63)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					v147 = v133
																					return v147
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
															v93 = m.Env.Pgmem_bn_op(m, int32(1), v89, v90, v63, v65, v91, v92, v75, v19)
															mBase = m.M
															if v93 < int32(0) {
																v116 = v79
																if v77 != 0 {
																	v120 = F___memset(m, v77, int32(0), v72)
																	mBase = m.M
																	F_pfree(m, v77)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		if v75 != 0 {
																			v124 = F___memset(m, v75, int32(0), v72)
																			mBase = m.M
																			F_pfree(m, v75)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int32(0)
																			} else {
																				if v73 == int32(0) {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				} else {
																					v130 = F___memset(m, v73, int32(0), v72)
																					mBase = m.M
																					F_pfree(m, v73)
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return int32(0)
																					} else {
																						v133 = v116
																						if v63 == int32(0) {
																							v147 = v133
																							return v147
																						} else {
																							v144 = F___memset(m, v63, int32(0), v62)
																							mBase = m.M
																							F_pfree(m, v63)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return int32(0)
																							} else {
																								v147 = v133
																								return v147
																							}
																						}
																					}
																				}
																			}
																		} else {
																			if v73 == int32(0) {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			} else {
																				v130 = F___memset(m, v73, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v73)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return int32(0)
																				} else {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	if v75 != 0 {
																		v124 = F___memset(m, v75, int32(0), v72)
																		mBase = m.M
																		F_pfree(m, v75)
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int32(0)
																		} else {
																			if v73 == int32(0) {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			} else {
																				v130 = F___memset(m, v73, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v73)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return int32(0)
																				} else {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				}
																			}
																		}
																	} else {
																		if v73 == int32(0) {
																			v133 = v116
																			if v63 == int32(0) {
																				v147 = v133
																				return v147
																			} else {
																				v144 = F___memset(m, v63, int32(0), v62)
																				mBase = m.M
																				F_pfree(m, v63)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					v147 = v133
																					return v147
																				}
																			}
																		} else {
																			v130 = F___memset(m, v73, int32(0), v72)
																			mBase = m.M
																			F_pfree(m, v73)
																			mBase = m.M
																			v132 = m.ExcPending
																			if v132 != 0 {
																				return int32(0)
																			} else {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
																v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
																v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
																v101 = m.Env.Pgmem_bn_op(m, int32(2), v97, v98, v75, v93, v99, v100, v77, v19)
																mBase = m.M
																if v101 < int32(0) {
																	v116 = v79
																	if v77 != 0 {
																		v120 = F___memset(m, v77, int32(0), v72)
																		mBase = m.M
																		F_pfree(m, v77)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return int32(0)
																		} else {
																			if v75 != 0 {
																				v124 = F___memset(m, v75, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v75)
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return int32(0)
																				} else {
																					if v73 == int32(0) {
																						v133 = v116
																						if v63 == int32(0) {
																							v147 = v133
																							return v147
																						} else {
																							v144 = F___memset(m, v63, int32(0), v62)
																							mBase = m.M
																							F_pfree(m, v63)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return int32(0)
																							} else {
																								v147 = v133
																								return v147
																							}
																						}
																					} else {
																						v130 = F___memset(m, v73, int32(0), v72)
																						mBase = m.M
																						F_pfree(m, v73)
																						mBase = m.M
																						v132 = m.ExcPending
																						if v132 != 0 {
																							return int32(0)
																						} else {
																							v133 = v116
																							if v63 == int32(0) {
																								v147 = v133
																								return v147
																							} else {
																								v144 = F___memset(m, v63, int32(0), v62)
																								mBase = m.M
																								F_pfree(m, v63)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return int32(0)
																								} else {
																									v147 = v133
																									return v147
																								}
																							}
																						}
																					}
																				}
																			} else {
																				if v73 == int32(0) {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				} else {
																					v130 = F___memset(m, v73, int32(0), v72)
																					mBase = m.M
																					F_pfree(m, v73)
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return int32(0)
																					} else {
																						v133 = v116
																						if v63 == int32(0) {
																							v147 = v133
																							return v147
																						} else {
																							v144 = F___memset(m, v63, int32(0), v62)
																							mBase = m.M
																							F_pfree(m, v63)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return int32(0)
																							} else {
																								v147 = v133
																								return v147
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		if v75 != 0 {
																			v124 = F___memset(m, v75, int32(0), v72)
																			mBase = m.M
																			F_pfree(m, v75)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int32(0)
																			} else {
																				if v73 == int32(0) {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				} else {
																					v130 = F___memset(m, v73, int32(0), v72)
																					mBase = m.M
																					F_pfree(m, v73)
																					mBase = m.M
																					v132 = m.ExcPending
																					if v132 != 0 {
																						return int32(0)
																					} else {
																						v133 = v116
																						if v63 == int32(0) {
																							v147 = v133
																							return v147
																						} else {
																							v144 = F___memset(m, v63, int32(0), v62)
																							mBase = m.M
																							F_pfree(m, v63)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return int32(0)
																							} else {
																								v147 = v133
																								return v147
																							}
																						}
																					}
																				}
																			}
																		} else {
																			if v73 == int32(0) {
																				v133 = v116
																				if v63 == int32(0) {
																					v147 = v133
																					return v147
																				} else {
																					v144 = F___memset(m, v63, int32(0), v62)
																					mBase = m.M
																					F_pfree(m, v63)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						v147 = v133
																						return v147
																					}
																				}
																			} else {
																				v130 = F___memset(m, v73, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v73)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return int32(0)
																				} else {
																					v133 = v116
																					if v63 == int32(0) {
																						v147 = v133
																						return v147
																					} else {
																						v144 = F___memset(m, v63, int32(0), v62)
																						mBase = m.M
																						F_pfree(m, v63)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							v147 = v133
																							return v147
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v104 = F_bytes_to_mpi(m, v73, v85)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
																		v107 = F_bytes_to_mpi(m, v77, v101)
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v107
																			if v107 != 0 {
																				v112 = int32(0)
																			} else {
																				v112 = int32(-109)
																			}
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
																			if v114 != 0 {
																				v115 = v112
																			} else {
																				v115 = int32(-109)
																			}
																			v116 = v115
																			if v77 != 0 {
																				v120 = F___memset(m, v77, int32(0), v72)
																				mBase = m.M
																				F_pfree(m, v77)
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return int32(0)
																				} else {
																					if v75 != 0 {
																						v124 = F___memset(m, v75, int32(0), v72)
																						mBase = m.M
																						F_pfree(m, v75)
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return int32(0)
																						} else {
																							if v73 == int32(0) {
																								v133 = v116
																								if v63 == int32(0) {
																									v147 = v133
																									return v147
																								} else {
																									v144 = F___memset(m, v63, int32(0), v62)
																									mBase = m.M
																									F_pfree(m, v63)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return int32(0)
																									} else {
																										v147 = v133
																										return v147
																									}
																								}
																							} else {
																								v130 = F___memset(m, v73, int32(0), v72)
																								mBase = m.M
																								F_pfree(m, v73)
																								mBase = m.M
																								v132 = m.ExcPending
																								if v132 != 0 {
																									return int32(0)
																								} else {
																									v133 = v116
																									if v63 == int32(0) {
																										v147 = v133
																										return v147
																									} else {
																										v144 = F___memset(m, v63, int32(0), v62)
																										mBase = m.M
																										F_pfree(m, v63)
																										mBase = m.M
																										v146 = m.ExcPending
																										if v146 != 0 {
																											return int32(0)
																										} else {
																											v147 = v133
																											return v147
																										}
																									}
																								}
																							}
																						}
																					} else {
																						if v73 == int32(0) {
																							v133 = v116
																							if v63 == int32(0) {
																								v147 = v133
																								return v147
																							} else {
																								v144 = F___memset(m, v63, int32(0), v62)
																								mBase = m.M
																								F_pfree(m, v63)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return int32(0)
																								} else {
																									v147 = v133
																									return v147
																								}
																							}
																						} else {
																							v130 = F___memset(m, v73, int32(0), v72)
																							mBase = m.M
																							F_pfree(m, v73)
																							mBase = m.M
																							v132 = m.ExcPending
																							if v132 != 0 {
																								return int32(0)
																							} else {
																								v133 = v116
																								if v63 == int32(0) {
																									v147 = v133
																									return v147
																								} else {
																									v144 = F___memset(m, v63, int32(0), v62)
																									mBase = m.M
																									F_pfree(m, v63)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return int32(0)
																									} else {
																										v147 = v133
																										return v147
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				if v75 != 0 {
																					v124 = F___memset(m, v75, int32(0), v72)
																					mBase = m.M
																					F_pfree(m, v75)
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return int32(0)
																					} else {
																						if v73 == int32(0) {
																							v133 = v116
																							if v63 == int32(0) {
																								v147 = v133
																								return v147
																							} else {
																								v144 = F___memset(m, v63, int32(0), v62)
																								mBase = m.M
																								F_pfree(m, v63)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return int32(0)
																								} else {
																									v147 = v133
																									return v147
																								}
																							}
																						} else {
																							v130 = F___memset(m, v73, int32(0), v72)
																							mBase = m.M
																							F_pfree(m, v73)
																							mBase = m.M
																							v132 = m.ExcPending
																							if v132 != 0 {
																								return int32(0)
																							} else {
																								v133 = v116
																								if v63 == int32(0) {
																									v147 = v133
																									return v147
																								} else {
																									v144 = F___memset(m, v63, int32(0), v62)
																									mBase = m.M
																									F_pfree(m, v63)
																									mBase = m.M
																									v146 = m.ExcPending
																									if v146 != 0 {
																										return int32(0)
																									} else {
																										v147 = v133
																										return v147
																									}
																								}
																							}
																						}
																					}
																				} else {
																					if v73 == int32(0) {
																						v133 = v116
																						if v63 == int32(0) {
																							v147 = v133
																							return v147
																						} else {
																							v144 = F___memset(m, v63, int32(0), v62)
																							mBase = m.M
																							F_pfree(m, v63)
																							mBase = m.M
																							v146 = m.ExcPending
																							if v146 != 0 {
																								return int32(0)
																							} else {
																								v147 = v133
																								return v147
																							}
																						}
																					} else {
																						v130 = F___memset(m, v73, int32(0), v72)
																						mBase = m.M
																						F_pfree(m, v73)
																						mBase = m.M
																						v132 = m.ExcPending
																						if v132 != 0 {
																							return int32(0)
																						} else {
																							v133 = v116
																							if v63 == int32(0) {
																								v147 = v133
																								return v147
																							} else {
																								v144 = F___memset(m, v63, int32(0), v62)
																								mBase = m.M
																								F_pfree(m, v63)
																								mBase = m.M
																								v146 = m.ExcPending
																								if v146 != 0 {
																									return int32(0)
																								} else {
																									v147 = v133
																									return v147
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
func F_pgp_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int64
	_ = v240
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v563 int32
	_ = v563
	var v567 int64
	_ = v567
	var v570 int64
	_ = v570
	var v573 int64
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	v9 = m.G0
	v11 = v9 - int32(288)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(288)
	return v656
L2:
	;
	v18 = F_pushf_create_mbuf_writer(m, v11+int32(12), l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v656 = int32(-13)
	goto L1
L5:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_pushf_free_all(m, v652)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L178
	}
L6:
	;
	return int32(0)
L7:
	;
	if v18 < int32(0) {
		v647 = v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v25 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v128 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v28
	goto L14
L13:
	;
	goto L14
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v31)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v30)
	v42 = v30 & int32(255)
	switch v42 {
	case 0:
		v110 = v42
		goto L16
	case 1:
		goto L21
	default:
		goto L19
	case 3:
		goto L20
	}
L15:
	;
	if v110 < int32(0) {
		v647 = v110
		goto L5
	} else {
		goto L35
	}
L16:
	;
	m.G0 = v37 + int32(16)
	goto L15
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v105)
	v110 = int32(0)
	goto L16
L18:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+15)))
	v105 = v96&int32(31) | int32(96)
	goto L17
L19:
	;
	v110 = int32(-121)
	goto L16
L20:
	;
	v50 = int32(-17)
	v54 = F_pg_strong_random(m, l0+int32(2), int32(8))
	mBase = m.M
	if v54 == int32(0) {
		v110 = v50
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v48 = F_pg_strong_random(m, l0+int32(2), int32(8))
	mBase = m.M
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v49 = int32(0)
	goto L24
L23:
	;
	v49 = int32(-17)
	goto L24
L24:
	;
	v110 = v49
	goto L16
L25:
	;
	v60 = F_pg_strong_random(m, v37+int32(15), int32(1))
	mBase = m.M
	if v60 == int32(0) {
		v110 = v50
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v32 == int32(-1) {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v67 = int32(0)
	goto L28
L28:
	;
	v79 = int32(base.Ui32(v67)>>(uint(int32(4))%32)) + int32(6)
	if base.Ui32(v32) <= base.Ui32((v67&int32(14)|int32(16))<<(uint(v79)%32)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v105 = int32(255)
	goto L17
L30:
	;
	v105 = v67
	goto L17
L31:
	;
	goto L32
L32:
	;
	v83 = v67 | int32(1)
	if base.Ui32(v32) <= base.Ui32((v83&int32(15)|int32(16))<<(uint(v79)%32)) {
		v105 = v83
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v91 = v67 + int32(2)
	if v91 != int32(256) {
		v67 = v91
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v123 = F_pgp_s2k_process(m, l0, v120, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if v123 < int32(0) {
		v647 = v123
		goto L5
	} else {
		goto L37
	}
L37:
	;
	goto L11
L38:
	;
	if v350 < int32(0) {
		v647 = v350
		goto L5
	} else {
		goto L94
	}
L39:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v226 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v228)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v230)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+19)) = uint8(v232)
	if v230 == int32(0) {
		v254 = v11 + int32(16) | v226
		goto L69
	} else {
		goto L70
	}
L40:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v215
	if v215 != 0 {
		goto L66
	} else {
		goto L67
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v131 == int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v135 = int32(0)
	v137 = v134 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v137) {
		v152 = v135
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v152
	v157 = int32(0)
	v161 = m.G0
	v163 = v161 - int32(16)
	m.G0 = v163
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v157
	v169 = F_open(m, int32(288304), v157, v163)
	mBase = m.M
	if v169 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	goto L45
L47:
	;
	if int32(base.Ui32(int32(487))>>(uint(v137)%32))&int32(1) == int32(0) {
		v152 = v135
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v137<<(uint(int32(2))%32))+uint32(_consts[1308])))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v152 = v151
	goto L46
L49:
	;
	if v202 == int32(0) {
		v647 = int32(-17)
		goto L5
	} else {
		goto L62
	}
L50:
	;
	v172 = int32(1)
	if v152 == int32(0) {
		v195 = v172
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v202 = v157
	goto L52
L52:
	;
	m.G0 = v163 + int32(16)
	goto L49
L53:
	;
	v197 = F_close(m, v169)
	mBase = m.M
	v202 = v195
	goto L52
L54:
	;
	v175 = l0 + int32(132)
	v176 = v152
	goto L55
L55:
	;
	v181 = F_read(m, v169, v175, v176)
	mBase = m.M
	if v181 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v195 = v172
	goto L53
L57:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v185 == int32(27) {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v190 = v176 - v181
	if v190 != 0 {
		v175 = v175 + v181
		v176 = v190
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v195 = int32(0)
	goto L53
L61:
	;
	goto L56
L62:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v209 == int32(0) {
		goto L39
	} else {
		goto L63
	}
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v213 = F_pgp_write_pubenc_sesskey(m, l0, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v350 = v213
	goto L38
L65:
	;
	goto L39
L66:
	;
	v221 = F__emscripten_memcpy_bulkmem(m, l0+int32(132), l0+int32(11), v215)
	mBase = m.M
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v255 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+20)) = v240
	if v230 != int32(3) {
		v254 = v11 + int32(16) | int32(12)
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v248)
	v254 = v11 + int32(16) | int32(13)
	goto L69
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v256)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	v263 = int32(0)
	v265 = F_pgp_cfb_create(m, v11+int32(276), v228, l0+int32(11), v262, v263, v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L75
	}
L73:
	;
	v293 = v254
	goto L74
L74:
	;
	v294 = int32(195)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+276)) = uint8(v294)
	v300 = v293 - (v11 + int32(16))
	if v300 <= int32(191) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	if v265 < int32(0) {
		v350 = v265
		goto L38
	} else {
		goto L76
	}
L76:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v273 = F_pgp_cfb_encrypt(m, v269, v11+int32(8), int32(1), v254)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v281 = F_pgp_cfb_encrypt(m, v275, l0+int32(132), v278, v254+int32(1))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	F_pgp_cfb_free(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v288 = v286 + int32(1)
	if v288 < int32(0) {
		v350 = v288
		goto L38
	} else {
		goto L80
	}
L80:
	;
	v293 = v288 + v254
	goto L74
L81:
	;
	v337 = F_pushf_write(m, v225, v11+int32(276), v333-(v11+int32(276)))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L88
	}
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v300)
	v333 = v11 + int32(278)
	goto L81
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(v300) <= base.Ui32(int32(8383)) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v309 = v300 - int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+278)) = uint8(v309)
	v314 = int32(base.Ui32(v309)>>(uint(int32(8))%32)) + int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v314)
	v333 = v11 + int32(279)
	goto L81
L86:
	;
	goto L87
L87:
	;
	v318 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v318)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+281)) = uint8(v300)
	v322 = int32(base.Ui32(v300) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+280)) = uint8(v322)
	v325 = int32(base.Ui32(v300) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+279)) = uint8(v325)
	v328 = int32(base.Ui32(v300) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+278)) = uint8(v328)
	v333 = v11 + int32(282)
	goto L81
L88:
	;
	if int32(0) <= v337 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v343 = F_pushf_write(m, v225, v11+int32(16), v300)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L92
	}
L90:
	;
	v345 = v337
	goto L91
L91:
	;
	v349 = F___memset(m, v11+int32(16), int32(0), v300)
	mBase = m.M
	goto L93
L92:
	;
	v345 = v343
	goto L91
L93:
	;
	v350 = v345
	goto L38
L94:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v359 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v360 = int32(-55)
	goto L97
L96:
	;
	v360 = int32(-46)
	goto L97
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v360)
	v365 = F_pushf_write(m, v356, v11+int32(16), int32(1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	if v365 < int32(0) {
		v647 = v365
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v372 = F_pushf_create(m, v11+int32(8), int32(4394668), l0, v356)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	if v372 < int32(0) {
		v647 = v372
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v376
	v381 = F_pushf_create(m, v11+int32(8), int32(4394684), l0, v376)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	if v381 < int32(0) {
		v647 = v381
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v387 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v393 = F_pushf_create(m, v11+int32(8), int32(4394700), l0, v385)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L6
	} else {
		goto L107
	}
L105:
	;
	v399 = v385
	goto L106
L106:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v403 = int32(0)
	v405 = v402 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v405) {
		v420 = v403
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v393 < int32(0) {
		v647 = v393
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v397
	v399 = v397
	goto L106
L109:
	;
	v421 = int32(0)
	v425 = m.G0
	v427 = v425 - int32(16)
	m.G0 = v427
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v421
	v433 = F_open(m, int32(288304), v421, v427)
	mBase = m.M
	if v433 != int32(-1) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	goto L109
L111:
	;
	if int32(base.Ui32(int32(487))>>(uint(v405)%32))&int32(1) == int32(0) {
		v420 = v403
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v405<<(uint(int32(2))%32))+uint32(_consts[1308])))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+16))
	v420 = v419
	goto L110
L113:
	;
	if v466 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v436 = int32(1)
	if v420 == int32(0) {
		v459 = v436
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v466 = v421
	goto L116
L116:
	;
	m.G0 = v427 + int32(16)
	goto L113
L117:
	;
	v461 = F_close(m, v433)
	mBase = m.M
	v466 = v459
	goto L116
L118:
	;
	v439 = v11 + int32(16)
	v440 = v420
	goto L119
L119:
	;
	v445 = F_read(m, v433, v439, v440)
	mBase = m.M
	if v445 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v459 = v436
	goto L117
L121:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v449 == int32(27) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v454 = v440 - v445
	if v454 != 0 {
		v439 = v439 + v445
		v440 = v454
		goto L119
	} else {
		goto L125
	}
L124:
	;
	v459 = int32(0)
	goto L117
L125:
	;
	goto L120
L126:
	;
	v647 = int32(-17)
	goto L5
L127:
	;
	goto L128
L128:
	;
	v475 = v11 + int32(16)
	v476 = v475 + v420
	v477 = int32(2)
	v479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v476-v477))))
	*(*uint16)(unsafe.Add(mBase, uint32(v476))) = uint16(v479)
	v484 = v420 + v477
	v485 = F_pushf_write(m, v399, v475, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	v490 = F___memset(m, v11+int32(16), int32(0), v484)
	mBase = m.M
	goto L130
L130:
	;
	if v485 < int32(0) {
		v647 = v485
		goto L5
	} else {
		goto L131
	}
L131:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v493 <= int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v554 != 0 {
		goto L150
	} else {
		goto L151
	}
L133:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v551 = v496
	goto L132
L134:
	;
	goto L135
L135:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v498 <= int32(0) {
		v551 = v497
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v503 = m.G0
	v505 = v503 - int32(16)
	m.G0 = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v505)+14)) = uint8(v507)
	v509 = int32(200)
	*(*uint8)(unsafe.Add(mBase, uint32(v505)+15)) = uint8(v509)
	v514 = F_pushf_write(m, v497, v505+int32(15), int32(1))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L138
	}
L137:
	;
	m.G0 = v505 + int32(16)
	if v542 < int32(0) {
		v647 = v542
		goto L5
	} else {
		goto L149
	}
L138:
	;
	if v514 < int32(0) {
		v542 = v514
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v521 = F_pushf_create(m, v505+int32(8), int32(4394668), l0, v497)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	if v521 < int32(0) {
		v542 = v521
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v505)+8))
	v529 = F_pushf_write(m, v525, v505+int32(14), int32(1))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	if int32(0) <= v529 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v505)+8))
	v534 = F_pgp_compress_filter(m, v11+int32(8), l0, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L146
	}
L144:
	;
	v538 = v529
	goto L145
L145:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v505)+8))
	F_pushf_free(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L6
	} else {
		goto L148
	}
L146:
	;
	if int32(0) <= v534 {
		v542 = v534
		goto L137
	} else {
		goto L147
	}
L147:
	;
	v538 = v534
	goto L145
L148:
	;
	v542 = v538
	goto L137
L149:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v548
	v551 = v548
	goto L132
L150:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v557 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v560 = int32(98)
	goto L152
L152:
	;
	v561 = F___time(m)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v561)
	v563 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v563)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v560)
	v567 = int64(base.Ui64(v561) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v567)
	v570 = int64(base.Ui64(v561) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+19)) = uint8(v570)
	v573 = int64(base.Ui64(v561) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v573)
	v575 = int32(203)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+287)) = uint8(v575)
	v580 = F_pushf_write(m, v551, v11+int32(287), int32(1))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L6
	} else {
		goto L156
	}
L153:
	;
	v558 = int32(117)
	goto L155
L154:
	;
	v558 = int32(116)
	goto L155
L155:
	;
	v560 = v558
	goto L152
L156:
	;
	if v580 < int32(0) {
		v647 = v580
		goto L5
	} else {
		goto L157
	}
L157:
	;
	v587 = F_pushf_create(m, v11+int32(276), int32(4394668), l0, v551)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	if v587 < int32(0) {
		v647 = v587
		goto L5
	} else {
		goto L159
	}
L159:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v595 = F_pushf_write(m, v591, v11+int32(16), int32(6))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	if v595 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	F_pushf_free(m, v597)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L6
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v597
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v604 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v647 = v595
	goto L5
L165:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v622 = v620 - v621
	goto L170
L166:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v607 == int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v613 = F_pushf_create(m, v11+int32(8), int32(4394716), l0, v597)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v613 < int32(0) {
		v647 = v613
		goto L5
	} else {
		goto L169
	}
L169:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v617
	goto L165
L170:
	;
	v627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v627)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v630
	v632 = v629 - v630
	if v622 < v632 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v639 = F_pushf_write(m, v637, v638, v634)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L175
	}
L172:
	;
	v634 = v622
	goto L174
L173:
	;
	v634 = v632
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v630 + v634
	goto L171
L175:
	;
	if v639 < int32(0) {
		v647 = v639
		goto L5
	} else {
		goto L176
	}
L176:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v644 = F_pushf_flush(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v647 = v644
	goto L5
L178:
	;
	v656 = v647
	goto L1
}
func F_pgp_get_keyid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v3
	v19 = F_pullf_create_mbuf_reader(m, v13+int32(24), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v292
L2:
	;
	return int32(0)
L3:
	;
	if v19 < int32(0) {
		v292 = v19
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = v3
	v30 = v3
	v32 = v3
	v33 = v3
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v41 = F_pgp_parse_pkt_hdr(m, v35, v13+int32(15), v13+int32(16), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	F_pullf_free(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L50
	}
L7:
	;
	goto L6
L8:
	;
	if v41 <= int32(0) {
		v152 = v41
		v154 = v29
		v155 = v30
		v157 = v32
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v50 = F_pgp_create_pkt_reader(m, v13+int32(20), v47, v48, v41, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v50 < int32(0) {
		v152 = v50
		v154 = v29
		v155 = v30
		v157 = v32
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v54 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	switch v55 - v54 {
	case 0:
		goto L16
	case 1, 9, 11, 12, 16, 60:
		v127 = v32
		goto L14
	case 2:
		goto L15
	default:
		goto L13
	case 4, 5:
		goto L18
	case 6, 13:
		goto L17
	case 8, 17:
		v134 = v50
		v135 = v54
		v136 = v29
		v137 = v30
		v139 = v32
		v140 = v33
		goto L12
	}
L12:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v142 != 0 {
		goto L45
	} else {
		goto L46
	}
L13:
	;
	v134 = int32(-100)
	v135 = int32(0)
	v136 = v29
	v137 = v30
	v139 = v32
	v140 = v33
	goto L12
L14:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v130 = F_pgp_skip_packet(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v127 = v32 + int32(1)
	goto L14
L16:
	;
	v102 = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v107 = F_pullf_read_fixed(m, v103, int32(1), v13+int32(28))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L35
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v70 = F__pgp_read_public_key(m, v65, v13+int32(28))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L25
	}
L18:
	;
	v58 = int32(0)
	if v33 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v134 = int32(-114)
	v135 = v58
	v136 = v29
	v137 = v30
	v139 = v32
	v140 = int32(1)
	goto L12
L20:
	;
	goto L21
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v62 = F_pgp_skip_packet(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v134 = v62
	v135 = v58
	v136 = v29
	v137 = v30
	v139 = v32
	v140 = int32(1)
	goto L12
L23:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v77)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v94
	F_pgp_key_free(m, v77)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L34
	}
L24:
	;
	F_pgp_key_free(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L33
	}
L25:
	;
	if v70 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v88 = v70
	v89 = v74
	goto L24
L27:
	;
	goto L28
L28:
	;
	v75 = F_pgp_skip_packet(m, v65)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v75 < int32(0) {
		v88 = v75
		v89 = v77
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v80 = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+5)))
	if base.Ui32(int32(16)) < base.Ui32(v81) {
		v88 = v80
		v89 = v77
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if int32(1)<<(uint(v81)%32)&int32(65542) != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v88 = v80
	v89 = v77
	goto L24
L33:
	;
	v134 = v88
	v135 = int32(0)
	v136 = v29
	v137 = v30
	v139 = v32
	v140 = v33
	goto L12
L34:
	;
	v98 = int32(1)
	v134 = v98
	v135 = int32(0)
	v136 = v29 + v98
	v137 = v30
	v139 = v32
	v140 = v33
	goto L12
L35:
	;
	v111 = base.B2i32(v107 < int32(0))
	if v107 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = v107
	goto L38
L37:
	;
	v112 = int32(-1)
	goto L38
L38:
	;
	v114 = v30 + int32(1)
	if v107 < int32(0) {
		v134 = v112
		v135 = v102
		v136 = v29
		v137 = v114
		v139 = v32
		v140 = v33
		goto L12
	} else {
		goto L39
	}
L39:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)))
	if v115 != int32(3) {
		v134 = v112
		v135 = v102
		v136 = v29
		v137 = v114
		v139 = v32
		v140 = v33
		goto L12
	} else {
		goto L40
	}
L40:
	;
	v119 = F_pullf_read_fixed(m, v103, int32(8), v13)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v119 < int32(0) {
		v134 = v119
		v135 = v102
		v136 = v29
		v137 = v114
		v139 = v32
		v140 = v33
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v123 = F_pgp_skip_packet(m, v103)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v134 = v123
	v135 = v102
	v136 = v29
	v137 = v114
	v139 = v32
	v140 = v33
	goto L12
L44:
	;
	v134 = v130
	v135 = int32(0)
	v136 = v29
	v137 = v30
	v139 = v127
	v140 = v33
	goto L12
L45:
	;
	F_pullf_free(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v145 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v145
	if (v135^int32(1))&base.B2i32(v145 <= v134) != 0 {
		v29 = v136
		v30 = v137
		v32 = v139
		v33 = v140
		goto L5
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v152 = v134
	v154 = v136
	v155 = v137
	v157 = v139
	goto L7
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v163 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_pullf_free(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v152 < int32(0) {
		v292 = v152
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v168 = int32(-114)
	if v155 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v171 = int32(-100)
	goto L58
L57:
	;
	v171 = v152
	goto L58
L58:
	;
	if v154 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v172 = v171
	goto L61
L60:
	;
	v172 = v152
	goto L61
L61:
	;
	if int32(1) < v154 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v175 = v168
	goto L64
L63:
	;
	v175 = v172
	goto L64
L64:
	;
	if int32(1) < v155 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v178 = v168
	goto L67
L66:
	;
	v178 = v175
	goto L67
L67:
	;
	if v178 < int32(0) {
		v292 = v178
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v154|v155 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v292 = int32(6)
	goto L1
L70:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if v182 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v157 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v185 = int32(509271)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[1309]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v186
	v188 = *(*int32)(unsafe.Add(mBase, _consts[1310]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+3)) = v188
	goto L69
L74:
	;
	goto L75
L75:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v192 = int32(15)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v195)
	v197 = int32(4)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v191)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v200)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)) = uint8(v206)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v202)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v211)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v217)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v213)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v222)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)) = uint8(v228)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v233)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v239)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v235)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v250)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v246)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v255)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v261)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v257)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v266)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)))
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v269)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268&v192)+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)) = uint8(v274)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v268)>>(uint(v197)%32)))+uint32(_consts[1311]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)) = uint8(v279)
	v292 = int32(16)
	goto L1
L76:
	;
	v292 = int32(-119)
	goto L1
L77:
	;
	goto L78
L78:
	;
	v285 = int32(509278)
	v286 = *(*int32)(unsafe.Add(mBase, _consts[1312]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v286
	v288 = *(*int32)(unsafe.Add(mBase, _consts[1313]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+3)) = v288
	goto L69
}
func F_pgp_load_digest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = int32(-100)
	v6 = l0 - int32(1)
	if base.Ui32(int32(9)) < base.Ui32(v6) {
		v27 = v4
		return v27
	} else {
		if int32(base.Ui32(int32(903))>>(uint(v6)%32))&int32(1) == int32(0) {
			v27 = v4
			return v27
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_consts[1315])))
			v22 = F_px_find_digest(m, v21, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 != 0 {
					v26 = int32(-104)
				} else {
					v26 = int32(0)
				}
				v27 = v26
				return v27
			}
		}
	}
}
func F_pgp_mpi_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if base.Ui32(int32(65536)) <= base.Ui32(l1) {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
		F_px_debug(m, int32(465898), v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v40 = int32(-100)
			m.G0 = v9 + int32(16)
			return v40
		}
	} else {
		v23 = int32(base.Ui32(l1+int32(7)) >> (uint(int32(3)) % 32))
		v26 = F_palloc(m, v23+int32(12))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l1
			v31 = v26 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = v31
			if v23 != 0 {
				v33 = F__emscripten_memcpy_bulkmem(m, v31, l0, v23)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
			v40 = int32(0)
			m.G0 = v9 + int32(16)
			return v40
		}
	}
}
func F_pgp_mpi_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v10 = int32(8)
	v14 = v9<<(uint(v10)%32) | int32(base.Ui32(v9)>>(uint(v10)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v14)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v19].(func(*base.Module, int32, int32, int32))(m, l0, v7+int32(14), int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		m.T0[v26].(func(*base.Module, int32, int32, int32))(m, l0, v24, v25)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return int32(0)
		}
	}
}
func F_pgp_mpi_read(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = F_pullf_read_fixed(m, l0, int32(2), v9+int32(14))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 < int32(0) {
			v51 = v14
			m.G0 = v9 + int32(16)
			return v51
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
			v24 = v20 | v21<<(uint(int32(8))%32)
			v28 = int32(base.Ui32(v24+int32(7)) >> (uint(int32(3)) % 32))
			v31 = F_palloc(m, v28+int32(12))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v24
				v36 = v31 + int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v36
				v38 = F_pullf_read_fixed(m, l0, v28, v36)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v38 < int32(0) {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
						v46 = F___memset(m, v31, int32(0), v43+int32(12))
						mBase = m.M
						F_pfree(m, v31)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v51 = v38
							m.G0 = v9 + int32(16)
							return v51
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
						v51 = v38
						m.G0 = v9 + int32(16)
						return v51
					}
				}
			}
		}
	}
}
func F_pgp_pub_decrypt_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v16 < int32(3) {
		v29 = int32(0)
		v30 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = F_decrypt_internal(m, int32(1), int32(0), v8, v13, v29, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v23 < int32(4) {
		v29 = v20
		v30 = int32(0)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v27 = F_pg_detoast_datum_packed(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = v20
	v30 = v27
	goto L4
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 != v8 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_pfree(m, v8)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v39 != v13 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	F_pfree(m, v13)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v43 < int32(3) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return v33
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v46 != v29 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v29)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v51 = v43
	goto L22
L22:
	;
	if base.I32_extend16_s(v51) < int32(4) {
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v51 = v50
	goto L22
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v30 == v55 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v30)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L18
}
func F_pgp_pub_decrypt_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v18 < int32(3) {
				v30 = v2
				v31 = v2
				v32 = int32(1)
				v34 = F_decrypt_internal(m, v32, v32, v11, v16, v31, v30)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = int32(1)
					v37 = v34 + v36
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
					v42 = v40 & v36
					if v42 != 0 {
						v43 = v37
					} else {
						v43 = v34 + int32(4)
					}
					if v40 == int32(1) {
						v46 = int32(4)
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
						if v48&int32(254) == int32(2) {
							v57 = v46
						} else {
							v57 = base.B2i32(v48 == int32(18)) << (uint(v46) % 32)
						}
						if v48 == int32(1) {
							v60 = v46
						} else {
							v60 = v57
						}
						v71 = v60
					} else {
						v61 = int32(1)
						if v42 != 0 {
							v71 = int32(base.Ui32(v40)>>(uint(v61)%32)) - v61
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
							v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v43, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v74 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v78 != v16 {
									F_pfree(m, v16)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v82 < int32(3) {
											return v34
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v85 != v31 {
												F_pfree(m, v31)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v90 = v89
													if base.I32_extend16_s(v90) < int32(4) {
														return v34
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v30 == v94 {
															return v34
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											} else {
												v90 = v82
												if base.I32_extend16_s(v90) < int32(4) {
													return v34
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v30 == v94 {
														return v34
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															return v34
														}
													}
												}
											}
										}
									}
								} else {
									v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v82 < int32(3) {
										return v34
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v85 != v31 {
											F_pfree(m, v31)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v90 = v89
												if base.I32_extend16_s(v90) < int32(4) {
													return v34
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v30 == v94 {
														return v34
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															return v34
														}
													}
												}
											}
										} else {
											v90 = v82
											if base.I32_extend16_s(v90) < int32(4) {
												return v34
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v30 == v94 {
													return v34
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														return v34
													}
												}
											}
										}
									}
								}
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v78 != v16 {
								F_pfree(m, v16)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v82 < int32(3) {
										return v34
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v85 != v31 {
											F_pfree(m, v31)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v90 = v89
												if base.I32_extend16_s(v90) < int32(4) {
													return v34
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v30 == v94 {
														return v34
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															return v34
														}
													}
												}
											}
										} else {
											v90 = v82
											if base.I32_extend16_s(v90) < int32(4) {
												return v34
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v30 == v94 {
													return v34
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														return v34
													}
												}
											}
										}
									}
								}
							} else {
								v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v82 < int32(3) {
									return v34
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v85 != v31 {
										F_pfree(m, v31)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
											v90 = v89
											if base.I32_extend16_s(v90) < int32(4) {
												return v34
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v30 == v94 {
													return v34
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														return v34
													}
												}
											}
										}
									} else {
										v90 = v82
										if base.I32_extend16_s(v90) < int32(4) {
											return v34
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v30 == v94 {
												return v34
											} else {
												F_pfree(m, v30)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													return v34
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v22 = F_pg_detoast_datum_packed(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
					if v24 < int32(4) {
						v30 = v2
						v31 = v22
						v32 = int32(1)
						v34 = F_decrypt_internal(m, v32, v32, v11, v16, v31, v30)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = int32(1)
							v37 = v34 + v36
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
							v42 = v40 & v36
							if v42 != 0 {
								v43 = v37
							} else {
								v43 = v34 + int32(4)
							}
							if v40 == int32(1) {
								v46 = int32(4)
								v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
								if v48&int32(254) == int32(2) {
									v57 = v46
								} else {
									v57 = base.B2i32(v48 == int32(18)) << (uint(v46) % 32)
								}
								if v48 == int32(1) {
									v60 = v46
								} else {
									v60 = v57
								}
								v71 = v60
							} else {
								v61 = int32(1)
								if v42 != 0 {
									v71 = int32(base.Ui32(v40)>>(uint(v61)%32)) - v61
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_pg_verifymbstr(m, v43, v71)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v74 != v11 {
									F_pfree(m, v11)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v78 != v16 {
											F_pfree(m, v16)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v82 < int32(3) {
													return v34
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v85 != v31 {
														F_pfree(m, v31)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v90 = v89
															if base.I32_extend16_s(v90) < int32(4) {
																return v34
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v30 == v94 {
																	return v34
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v90 = v82
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												}
											}
										} else {
											v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v82 < int32(3) {
												return v34
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v85 != v31 {
													F_pfree(m, v31)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v90 = v89
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v90 = v82
													if base.I32_extend16_s(v90) < int32(4) {
														return v34
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v30 == v94 {
															return v34
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											}
										}
									}
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v78 != v16 {
										F_pfree(m, v16)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v82 < int32(3) {
												return v34
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v85 != v31 {
													F_pfree(m, v31)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v90 = v89
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v90 = v82
													if base.I32_extend16_s(v90) < int32(4) {
														return v34
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v30 == v94 {
															return v34
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											}
										}
									} else {
										v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v82 < int32(3) {
											return v34
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v85 != v31 {
												F_pfree(m, v31)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v90 = v89
													if base.I32_extend16_s(v90) < int32(4) {
														return v34
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v30 == v94 {
															return v34
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											} else {
												v90 = v82
												if base.I32_extend16_s(v90) < int32(4) {
													return v34
												} else {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v30 == v94 {
														return v34
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															return v34
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
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v28 = F_pg_detoast_datum_packed(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = v28
							v31 = v22
							v32 = int32(1)
							v34 = F_decrypt_internal(m, v32, v32, v11, v16, v31, v30)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = int32(1)
								v37 = v34 + v36
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
								v42 = v40 & v36
								if v42 != 0 {
									v43 = v37
								} else {
									v43 = v34 + int32(4)
								}
								if v40 == int32(1) {
									v46 = int32(4)
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
									if v48&int32(254) == int32(2) {
										v57 = v46
									} else {
										v57 = base.B2i32(v48 == int32(18)) << (uint(v46) % 32)
									}
									if v48 == int32(1) {
										v60 = v46
									} else {
										v60 = v57
									}
									v71 = v60
								} else {
									v61 = int32(1)
									if v42 != 0 {
										v71 = int32(base.Ui32(v40)>>(uint(v61)%32)) - v61
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
										v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_pg_verifymbstr(m, v43, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v74 != v11 {
										F_pfree(m, v11)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v78 != v16 {
												F_pfree(m, v16)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v82 < int32(3) {
														return v34
													} else {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v85 != v31 {
															F_pfree(m, v31)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
																v90 = v89
																if base.I32_extend16_s(v90) < int32(4) {
																	return v34
																} else {
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																	if v30 == v94 {
																		return v34
																	} else {
																		F_pfree(m, v30)
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return int32(0)
																		} else {
																			return v34
																		}
																	}
																}
															}
														} else {
															v90 = v82
															if base.I32_extend16_s(v90) < int32(4) {
																return v34
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v30 == v94 {
																	return v34
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													}
												}
											} else {
												v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v82 < int32(3) {
													return v34
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v85 != v31 {
														F_pfree(m, v31)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v90 = v89
															if base.I32_extend16_s(v90) < int32(4) {
																return v34
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v30 == v94 {
																	return v34
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v90 = v82
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												}
											}
										}
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v78 != v16 {
											F_pfree(m, v16)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v82 < int32(3) {
													return v34
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v85 != v31 {
														F_pfree(m, v31)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v90 = v89
															if base.I32_extend16_s(v90) < int32(4) {
																return v34
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v30 == v94 {
																	return v34
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v90 = v82
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												}
											}
										} else {
											v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v82 < int32(3) {
												return v34
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v85 != v31 {
													F_pfree(m, v31)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v90 = v89
														if base.I32_extend16_s(v90) < int32(4) {
															return v34
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v30 == v94 {
																return v34
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v90 = v82
													if base.I32_extend16_s(v90) < int32(4) {
														return v34
													} else {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v30 == v94 {
															return v34
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																return v34
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
}
func F_pgp_pub_encrypt_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v14 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v23 = F_encrypt_internal(m, int32(1), int32(0), v7, v12, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v7 {
							F_pfree(m, v7)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v12 {
									F_pfree(m, v12)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v20 == v36 {
												return v23
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(0)
				v23 = F_encrypt_internal(m, int32(1), int32(0), v7, v12, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v20 == v36 {
									return v23
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
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
func F_pgp_pub_encrypt_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v14 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v21 = int32(1)
					v23 = F_encrypt_internal(m, v21, v21, v7, v12, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v7 {
							F_pfree(m, v7)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v12 {
									F_pfree(m, v12)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v20 == v36 {
												return v23
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(0)
				v21 = int32(1)
				v23 = F_encrypt_internal(m, v21, v21, v7, v12, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v20 == v36 {
									return v23
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
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
func F_pgp_set_pubkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(704)
	m.G0 = v20
	v24 = F_pullf_create_mbuf_reader(m, v20+int32(36), l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(704)
	return v853
L2:
	;
	return int32(0)
L3:
	;
	if v24 < int32(0) {
		v853 = v24
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v31
	v47 = v6
	v48 = v6
	v51 = v6
	goto L5
L5:
	;
	v59 = F_pgp_parse_pkt_hdr(m, v30, v20+int32(51), v20+int32(44), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	if v833 != 0 {
		goto L218
	} else {
		goto L219
	}
L7:
	;
	goto L6
L8:
	;
	if v59 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v823 = v59
	v826 = v48
	goto L7
L10:
	;
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v67 = F_pgp_create_pkt_reader(m, v20+int32(52), v30, v65, v59, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v67 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v823 = v67
	v826 = v48
	goto L7
L14:
	;
	goto L15
L15:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)))
	switch v71 - int32(2) {
	case 0, 8, 10, 11, 15, 59:
		goto L20
	default:
		goto L19
	case 3, 4:
		goto L23
	case 5:
		goto L21
	case 12:
		goto L22
	}
L16:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	F_pullf_free(m, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L2
	} else {
		goto L199
	}
L17:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
	switch v215 - int32(1) {
	case 0, 1, 2:
		goto L72
	default:
		goto L70
	case 15, 16:
		v237 = int32(24)
		goto L71
	}
L18:
	;
	v776 = v204
	v779 = int32(0)
	v781 = v47
	v783 = v51
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v71
	F_px_debug(m, int32(481341), v20)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L68
	}
L20:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v197 = F_pgp_skip_packet(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L67
	}
L21:
	;
	if l4 != int32(1) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	if l4 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v74 = int32(0)
	if v51 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v776 = int32(-114)
	v779 = v74
	v781 = v47
	v783 = int32(1)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v78 = F_pgp_skip_packet(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v776 = v78
	v779 = v74
	v781 = v47
	v783 = int32(1)
	goto L16
L28:
	;
	v204 = int32(-116)
	goto L18
L29:
	;
	goto L30
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v85 = F__pgp_read_public_key(m, v82, v20+int32(40))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v776 = v85
	v779 = v87
	v781 = v87
	v783 = v51
	goto L16
L32:
	;
	v204 = int32(-115)
	goto L18
L33:
	;
	goto L34
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v92
	v99 = F__pgp_read_public_key(m, v91, v20+int32(56))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v99 < int32(0) {
		v776 = v99
		v779 = v92
		v781 = v47
		v783 = v51
		goto L16
	} else {
		goto L36
	}
L36:
	;
	v106 = F_pullf_read_fixed(m, v91, int32(1), v20+int32(112))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v106 < int32(0) {
		v776 = v106
		v779 = v92
		v781 = v47
		v783 = v51
		goto L16
	} else {
		goto L38
	}
L38:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+112)))
	if base.Ui32(int32(254)) <= base.Ui32(v110) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l2 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v110 == int32(0) {
		v212 = v91
		goto L17
	} else {
		goto L65
	}
L42:
	;
	v776 = int32(-120)
	v779 = v92
	v781 = v47
	v783 = v51
	goto L16
L43:
	;
	goto L44
L44:
	;
	v119 = F_pullf_read_fixed(m, v91, int32(1), v20+int32(112))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	if v119 < int32(0) {
		v204 = v119
		goto L18
	} else {
		goto L46
	}
L46:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+112)))
	v126 = F_pgp_s2k_read(m, v91, v20+int32(60))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	if v126 < int32(0) {
		v204 = v126
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v132 = F_pgp_s2k_process(m, v20+int32(60), v123, l2, l3)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v132 < int32(0) {
		v204 = v132
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v136 = int32(0)
	v138 = v123 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v138) {
		v153 = v136
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v153 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	if int32(base.Ui32(int32(487))>>(uint(v138)%32))&int32(1) == int32(0) {
		v153 = v136
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138<<(uint(int32(2))%32))+uint32(_consts[1308])))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v153 = v152
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v123
	F_px_debug(m, int32(466009), v20+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v165 = F_pullf_read_fixed(m, v91, v153, v20+int32(112))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L59
	}
L58:
	;
	v204 = int32(-103)
	goto L18
L59:
	;
	if v165 < int32(0) {
		v204 = v165
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+103)))
	v175 = F_pgp_cfb_create(m, v20+int32(104), v123, v20+int32(71), v171, int32(0), v20+int32(112))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v175 < int32(0) {
		v204 = v175
		goto L18
	} else {
		goto L62
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v183 = F_pullf_create(m, v20+int32(108), int32(4394620), v182, v91)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	if v183 < int32(0) {
		v204 = v183
		goto L18
	} else {
		goto L64
	}
L64:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v212 = v187
	goto L17
L65:
	;
	F_px_debug(m, int32(370231), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v776 = int32(-118)
	v779 = v92
	v781 = v47
	v783 = v51
	goto L16
L67:
	;
	v776 = v197
	v779 = int32(0)
	v781 = v47
	v783 = v51
	goto L16
L68:
	;
	v204 = int32(-107)
	goto L18
L69:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v765 != 0 {
		goto L187
	} else {
		goto L188
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v215
	F_px_debug(m, int32(482589), v20+int32(32))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L2
	} else {
		goto L186
	}
L71:
	;
	v239 = F_pgp_mpi_read(m, v212, v237+v214)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L79
	}
L72:
	;
	v220 = F_pgp_mpi_read(m, v212, v214+int32(24))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v220 < int32(0) {
		v761 = v220
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v226 = F_pgp_mpi_read(m, v212, v214+int32(28))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if v226 < int32(0) {
		v761 = v226
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v232 = F_pgp_mpi_read(m, v212, v214+int32(32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v232 < int32(0) {
		v761 = v232
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v237 = int32(36)
	goto L71
L79:
	;
	if v239 < int32(0) {
		v761 = v239
		goto L69
	} else {
		goto L80
	}
L80:
	;
	if v110 == int32(254) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v746 < int32(0) {
		v761 = v746
		goto L69
	} else {
		goto L184
	}
L82:
	;
	v248 = F_pullf_read_fixed(m, v212, int32(20), v20+int32(672))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v382 = F_pullf_read_fixed(m, v212, int32(2), v20+int32(672))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L121
	}
L85:
	;
	if v248 < int32(0) {
		v746 = v248
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v255 = F_pgp_load_digest(m, int32(2), v20+int32(636))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L88
	}
L87:
	;
	v373 = F___memset(m, v20+int32(672), int32(0), int32(20))
	mBase = m.M
	goto L119
L88:
	;
	if v255 < int32(0) {
		v367 = v255
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
	switch v260 - int32(1) {
	case 0, 1, 2:
		goto L92
	default:
		goto L90
	case 15, 16:
		v276 = int32(24)
		goto L91
	}
L90:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
	m.T0[v286].(func(*base.Module, int32, int32))(m, v283, v20+int32(640))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L97
	}
L91:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v214+v276)))
	v280 = F_pgp_mpi_hash(m, v277, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L96
	}
L92:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	v265 = F_pgp_mpi_hash(m, v263, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v214)+28))
	v269 = F_pgp_mpi_hash(m, v267, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v214)+32))
	v273 = F_pgp_mpi_hash(m, v271, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v276 = int32(36)
	goto L91
L96:
	;
	goto L90
L97:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
	m.T0[v290].(func(*base.Module, int32))(m, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v294 = v20 + int32(640)
	v296 = v20 + int32(672)
	v297 = int32(20)
	goto L102
L99:
	;
	if v359 == int32(0) {
		v367 = v255
		goto L87
	} else {
		goto L117
	}
L100:
	;
	v359 = int32(0)
	goto L99
L101:
	;
	v333 = v328
	v334 = v329
	v335 = v330
	goto L111
L102:
	;
	if (v294|v296)&int32(3) != 0 {
		v328 = v294
		v329 = v296
		v330 = v297
		goto L101
	} else {
		goto L105
	}
L104:
	;
	if v318 == int32(0) {
		goto L100
	} else {
		goto L110
	}
L105:
	;
	v305 = v294
	v306 = v296
	v307 = v297
	goto L106
L106:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v310 != v311 {
		v328 = v305
		v329 = v306
		v330 = v307
		goto L101
	} else {
		goto L108
	}
L107:
	;
	goto L104
L108:
	;
	v313 = int32(4)
	v314 = v306 + v313
	v316 = v305 + v313
	v318 = v307 - v313
	if base.Ui32(int32(3)) < base.Ui32(v318) {
		v305 = v316
		v306 = v314
		v307 = v318
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v328 = v316
	v329 = v314
	v330 = v318
	goto L101
L111:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v338 == v339 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v359 = v338 - v339
	goto L99
L113:
	;
	v341 = int32(1)
	v346 = v335 - v341
	if v346 != 0 {
		v333 = v333 + v341
		v334 = v334 + v341
		v335 = v346
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	goto L100
L117:
	;
	F_px_debug(m, int32(454012), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v367 = int32(-118)
	goto L87
L119:
	;
	v378 = F___memset(m, v20+int32(640), int32(0), int32(20))
	mBase = m.M
	goto L120
L120:
	;
	v746 = v367
	goto L81
L121:
	;
	if v382 < int32(0) {
		v746 = v382
		goto L81
	} else {
		goto L122
	}
L122:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+673)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+672)))
	v392 = int32(0)
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
	switch v394 - int32(1) {
	case 0, 1, 2:
		goto L125
	default:
		v738 = v392
		goto L123
	case 15, 16:
		v651 = v392
		v652 = int32(24)
		goto L124
	}
L123:
	;
	if v738 == v386|v387<<(uint(int32(8))%32) {
		v746 = v392
		goto L81
	} else {
		goto L182
	}
L124:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v214+v652)))
	v655 = int32(0)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v654)+4))
	v667 = v661>>(uint(int32(8))%32) + v651 + v661&int32(255)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v654)+8))
	if v668 <= v655 {
		v728 = v667
		goto L169
	} else {
		goto L170
	}
L125:
	;
	v398 = int32(0)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v412 = v406>>(uint(int32(8))%32) + v398 + v406&int32(255)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	if v413 <= v398 {
		v473 = v412
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v214)+28))
	v484 = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	v496 = v490>>(uint(int32(8))%32) + v473&int32(65535) + v490&int32(255)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	if v497 <= v484 {
		v557 = v496
		goto L141
	} else {
		goto L142
	}
L127:
	;
	goto L126
L128:
	;
	v417 = v413 & int32(3)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	if base.Ui32(v413) < base.Ui32(int32(4)) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v417 == int32(0) {
		v473 = v447
		goto L127
	} else {
		goto L136
	}
L130:
	;
	v447 = v412
	v448 = int32(0)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v425 = v412
	v426 = int32(0)
	v431 = v398
	goto L133
L133:
	;
	v433 = v426 + v418
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+2)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+3)))
	v441 = v425 + v434 + v436 + v438 + v440
	v442 = int32(4)
	v443 = v426 + v442
	v445 = v431 + v442
	if v445 != v413&int32(2147483644) {
		v425 = v441
		v426 = v443
		v431 = v445
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v447 = v441
	v448 = v443
	goto L129
L135:
	;
	goto L134
L136:
	;
	v457 = v447
	v458 = v448
	v462 = v398
	goto L137
L137:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v418))))
	v467 = v457 + v466
	v468 = int32(1)
	v471 = v462 + v468
	if v471 != v417 {
		v457 = v467
		v458 = v458 + v468
		v462 = v471
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v473 = v467
	goto L127
L139:
	;
	goto L138
L140:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v214)+32))
	v568 = int32(0)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v580 = v574>>(uint(int32(8))%32) + v557&int32(65535) + v574&int32(255)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	if v581 <= v568 {
		v641 = v580
		goto L155
	} else {
		goto L156
	}
L141:
	;
	goto L140
L142:
	;
	v501 = v497 & int32(3)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	if base.Ui32(v497) < base.Ui32(int32(4)) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v501 == int32(0) {
		v557 = v531
		goto L141
	} else {
		goto L150
	}
L144:
	;
	v531 = v496
	v532 = int32(0)
	goto L143
L145:
	;
	goto L146
L146:
	;
	v509 = v496
	v510 = int32(0)
	v515 = v484
	goto L147
L147:
	;
	v517 = v510 + v502
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517))))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+2)))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+3)))
	v525 = v509 + v518 + v520 + v522 + v524
	v526 = int32(4)
	v527 = v510 + v526
	v529 = v515 + v526
	if v529 != v497&int32(2147483644) {
		v509 = v525
		v510 = v527
		v515 = v529
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v531 = v525
	v532 = v527
	goto L143
L149:
	;
	goto L148
L150:
	;
	v541 = v531
	v542 = v532
	v546 = v484
	goto L151
L151:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+v502))))
	v551 = v541 + v550
	v552 = int32(1)
	v555 = v546 + v552
	if v555 != v501 {
		v541 = v551
		v542 = v542 + v552
		v546 = v555
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v557 = v551
	goto L141
L153:
	;
	goto L152
L154:
	;
	v651 = v641 & int32(65535)
	v652 = int32(36)
	goto L124
L155:
	;
	goto L154
L156:
	;
	v585 = v581 & int32(3)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if base.Ui32(v581) < base.Ui32(int32(4)) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v585 == int32(0) {
		v641 = v615
		goto L155
	} else {
		goto L164
	}
L158:
	;
	v615 = v580
	v616 = int32(0)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v593 = v580
	v594 = int32(0)
	v599 = v568
	goto L161
L161:
	;
	v601 = v594 + v586
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+2)))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+3)))
	v609 = v593 + v602 + v604 + v606 + v608
	v610 = int32(4)
	v611 = v594 + v610
	v613 = v599 + v610
	if v613 != v581&int32(2147483644) {
		v593 = v609
		v594 = v611
		v599 = v613
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v615 = v609
	v616 = v611
	goto L157
L163:
	;
	goto L162
L164:
	;
	v625 = v615
	v626 = v616
	v630 = v568
	goto L165
L165:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v586))))
	v635 = v625 + v634
	v636 = int32(1)
	v639 = v630 + v636
	if v639 != v585 {
		v625 = v635
		v626 = v626 + v636
		v630 = v639
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v641 = v635
	goto L155
L167:
	;
	goto L166
L168:
	;
	v738 = v728 & int32(65535)
	goto L123
L169:
	;
	goto L168
L170:
	;
	v672 = v668 & int32(3)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	if base.Ui32(v668) < base.Ui32(int32(4)) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v672 == int32(0) {
		v728 = v702
		goto L169
	} else {
		goto L178
	}
L172:
	;
	v702 = v667
	v703 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v680 = v667
	v681 = int32(0)
	v686 = v655
	goto L175
L175:
	;
	v688 = v681 + v673
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+1)))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+2)))
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+3)))
	v696 = v680 + v689 + v691 + v693 + v695
	v697 = int32(4)
	v698 = v681 + v697
	v700 = v686 + v697
	if v700 != v668&int32(2147483644) {
		v680 = v696
		v681 = v698
		v686 = v700
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v702 = v696
	v703 = v698
	goto L171
L177:
	;
	goto L176
L178:
	;
	v712 = v702
	v713 = v703
	v717 = v655
	goto L179
L179:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713+v673))))
	v722 = v712 + v721
	v723 = int32(1)
	v726 = v717 + v723
	if v726 != v672 {
		v712 = v722
		v713 = v713 + v723
		v717 = v726
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v728 = v722
	goto L169
L181:
	;
	goto L180
L182:
	;
	F_px_debug(m, int32(453989), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L2
	} else {
		goto L183
	}
L183:
	;
	v746 = int32(-118)
	goto L81
L184:
	;
	v752 = F_pgp_expect_packet_end(m, v212)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	v761 = v752
	goto L69
L186:
	;
	v761 = int32(-118)
	goto L69
L187:
	;
	F_pullf_free(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L2
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v768 != 0 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L189
L191:
	;
	F_pgp_cfb_free(m, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L2
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	if v761 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	F_pgp_key_free(m, v214)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L2
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v214
	v776 = v761
	v779 = v214
	v781 = v214
	v783 = v51
	goto L16
L198:
	;
	v776 = v761
	v779 = v92
	v781 = v47
	v783 = v51
	goto L16
L199:
	;
	v788 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v788
	if v779 == v788 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if int32(0) <= v816 {
		v47 = v819
		v48 = v820
		v51 = v783
		goto L5
	} else {
		goto L217
	}
L201:
	;
	v816 = v776
	v818 = v48
	v819 = v781
	v820 = v48
	goto L200
L202:
	;
	goto L203
L203:
	;
	if v776 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v813 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v813
	v816 = v809
	v818 = v811
	v819 = v813
	v820 = v811
	goto L200
L205:
	;
	v809 = v807
	v811 = v48
	goto L204
L206:
	;
	F_pgp_key_free(m, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L2
	} else {
		goto L216
	}
L207:
	;
	v803 = v776
	v804 = v779
	goto L206
L208:
	;
	goto L209
L209:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v779)+48))
	if v794 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v781 == int32(0) {
		v807 = v800
		goto L205
	} else {
		goto L215
	}
L211:
	;
	v800 = v776
	goto L210
L212:
	;
	goto L213
L213:
	;
	if v48 == int32(0) {
		v809 = v776
		v811 = v779
		goto L204
	} else {
		goto L214
	}
L214:
	;
	v800 = int32(-123)
	goto L210
L215:
	;
	v803 = v800
	v804 = v781
	goto L206
L216:
	;
	v807 = v803
	goto L205
L217:
	;
	v823 = v816
	v826 = v818
	goto L7
L218:
	;
	F_pullf_free(m, v833)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L2
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if v823 < int32(0) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	goto L220
L222:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L2
	} else {
		goto L231
	}
L223:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L2
	} else {
		goto L230
	}
L224:
	;
	if v826 == int32(0) {
		v843 = v823
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	if v826 != 0 {
		goto L222
	} else {
		goto L229
	}
L227:
	;
	F_pgp_key_free(m, v826)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v843 = v823
	goto L223
L229:
	;
	v843 = int32(-119)
	goto L223
L230:
	;
	v853 = v843
	goto L1
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v826
	v853 = int32(0)
	goto L1
}
func F_pgp_set_s2k_count(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = int32(-13)
	if base.Ui32(int32(65010688)) < base.Ui32(l1-int32(1024)) {
		v14 = v4
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v9 != int32(3) {
			v14 = v4
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = l1
			v14 = int32(0)
		}
	}
	return v14
}
func F_pgp_set_sess_key(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = base.B2i32(l1 != v3)
	return v3
}
func F_pgp_sym_decrypt_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v17 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v21 = F_pg_detoast_datum_packed(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = v21
					v24 = int32(0)
					v27 = F_decrypt_internal(m, v24, int32(1), v10, v15, v24, v23)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(1)
						v30 = v27 + v29
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
						v35 = v33 & v29
						if v35 != 0 {
							v36 = v30
						} else {
							v36 = v27 + int32(4)
						}
						if v33 == int32(1) {
							v39 = int32(4)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
							if v41&int32(254) == int32(2) {
								v50 = v39
							} else {
								v50 = base.B2i32(v41 == int32(18)) << (uint(v39) % 32)
							}
							if v41 == int32(1) {
								v53 = v39
							} else {
								v53 = v50
							}
							v64 = v53
						} else {
							v54 = int32(1)
							if v35 != 0 {
								v64 = int32(base.Ui32(v33)>>(uint(v54)%32)) - v54
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								v64 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_pg_verifymbstr(m, v36, v64)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v67 != v10 {
								F_pfree(m, v10)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v71 != v15 {
										F_pfree(m, v15)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v75 < int32(3) {
												return v27
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v23 == v78 {
													return v27
												} else {
													F_pfree(m, v23)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v27
													}
												}
											}
										}
									} else {
										v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v75 < int32(3) {
											return v27
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v23 == v78 {
												return v27
											} else {
												F_pfree(m, v23)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								}
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v71 != v15 {
									F_pfree(m, v15)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v75 < int32(3) {
											return v27
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v23 == v78 {
												return v27
											} else {
												F_pfree(m, v23)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								} else {
									v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v75 < int32(3) {
										return v27
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v23 == v78 {
											return v27
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v27
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v23 = int32(0)
				v24 = int32(0)
				v27 = F_decrypt_internal(m, v24, int32(1), v10, v15, v24, v23)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(1)
					v30 = v27 + v29
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					v35 = v33 & v29
					if v35 != 0 {
						v36 = v30
					} else {
						v36 = v27 + int32(4)
					}
					if v33 == int32(1) {
						v39 = int32(4)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
						if v41&int32(254) == int32(2) {
							v50 = v39
						} else {
							v50 = base.B2i32(v41 == int32(18)) << (uint(v39) % 32)
						}
						if v41 == int32(1) {
							v53 = v39
						} else {
							v53 = v50
						}
						v64 = v53
					} else {
						v54 = int32(1)
						if v35 != 0 {
							v64 = int32(base.Ui32(v33)>>(uint(v54)%32)) - v54
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							v64 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v36, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v67 != v10 {
							F_pfree(m, v10)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v71 != v15 {
									F_pfree(m, v15)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v75 < int32(3) {
											return v27
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v23 == v78 {
												return v27
											} else {
												F_pfree(m, v23)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								} else {
									v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v75 < int32(3) {
										return v27
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v23 == v78 {
											return v27
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v27
											}
										}
									}
								}
							}
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v71 != v15 {
								F_pfree(m, v15)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v75 < int32(3) {
										return v27
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v23 == v78 {
											return v27
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v27
											}
										}
									}
								}
							} else {
								v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v75 < int32(3) {
									return v27
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v23 == v78 {
										return v27
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v27
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
func F_pgp_sym_encrypt_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v14 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = F_pg_detoast_datum_packed(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					v21 = int32(0)
					v23 = F_encrypt_internal(m, v21, v21, v7, v12, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v7 {
							F_pfree(m, v7)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v12 {
									F_pfree(m, v12)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v20 == v36 {
												return v23
											} else {
												F_pfree(m, v20)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v20 = int32(0)
				v21 = int32(0)
				v23 = F_encrypt_internal(m, v21, v21, v7, v12, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v20 == v36 {
											return v23
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v20 == v36 {
										return v23
									} else {
										F_pfree(m, v20)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v20 == v36 {
									return v23
								} else {
									F_pfree(m, v20)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
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
func F_pgp_write_pubenc_sesskey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v11 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v3
	if v10 == v3 {
		F_px_debug(m, int32(752131), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v154 = int32(-12)
			m.G0 = v8 + int32(32)
			return v154
		}
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v24)
		v29 = F_pgp_create_pkt_writer(m, l1, int32(1), v8+int32(12))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 < int32(0) {
				v146 = v29
				v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if v148 == int32(0) {
					v154 = v146
					m.G0 = v8 + int32(32)
					return v154
				} else {
					F_pushf_free(m, v148)
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return int32(0)
					} else {
						v154 = v146
						m.G0 = v8 + int32(32)
						return v154
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v37 = F_pushf_write(m, v33, v8+int32(19), int32(1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 < int32(0) {
						v146 = v37
						v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						if v148 == int32(0) {
							v154 = v146
							m.G0 = v8 + int32(32)
							return v154
						} else {
							F_pushf_free(m, v148)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								v154 = v146
								m.G0 = v8 + int32(32)
								return v154
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v45 = F_pushf_write(m, v41, v10+int32(40), int32(8))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 < int32(0) {
								v146 = v45
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								if v148 == int32(0) {
									v154 = v146
									m.G0 = v8 + int32(32)
									return v154
								} else {
									F_pushf_free(m, v148)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v154 = v146
										m.G0 = v8 + int32(32)
										return v154
									}
								}
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v53 = F_pushf_write(m, v49, v8+int32(11), int32(1))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53 < int32(0) {
										v146 = v53
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
										if v148 == int32(0) {
											v154 = v146
											m.G0 = v8 + int32(32)
											return v154
										} else {
											F_pushf_free(m, v148)
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int32(0)
											} else {
												v154 = v146
												m.G0 = v8 + int32(32)
												return v154
											}
										}
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
										switch v57 - int32(1) {
										case 0, 1:
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v104 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v104
											*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v104
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
											v114 = F_create_secmsg(m, l0, v8+int32(28), v111-int32(1))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												if v114 < int32(0) {
													v128 = v114
													v129 = F_pgp_mpi_free(m, v116)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
														v132 = F_pgp_mpi_free(m, v131)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															v135 = v128
															if v135 < int32(0) {
																v146 = v135
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																if v148 == int32(0) {
																	v154 = v146
																	m.G0 = v8 + int32(32)
																	return v154
																} else {
																	F_pushf_free(m, v148)
																	mBase = m.M
																	v152 = m.ExcPending
																	if v152 != 0 {
																		return int32(0)
																	} else {
																		v154 = v146
																		m.G0 = v8 + int32(32)
																		return v154
																	}
																}
															} else {
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																v143 = F_pushf_flush(m, v142)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return int32(0)
																} else {
																	v146 = v143
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																	if v148 == int32(0) {
																		v154 = v146
																		m.G0 = v8 + int32(32)
																		return v154
																	} else {
																		F_pushf_free(m, v148)
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			v154 = v146
																			m.G0 = v8 + int32(32)
																			return v154
																		}
																	}
																}
															}
														}
													}
												} else {
													v121 = F_pgp_rsa_encrypt(m, v10, v116, v8+int32(24))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														if v121 < int32(0) {
															v128 = v121
															v129 = F_pgp_mpi_free(m, v116)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return int32(0)
															} else {
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
																v132 = F_pgp_mpi_free(m, v131)
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return int32(0)
																} else {
																	v135 = v128
																	if v135 < int32(0) {
																		v146 = v135
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																		if v148 == int32(0) {
																			v154 = v146
																			m.G0 = v8 + int32(32)
																			return v154
																		} else {
																			F_pushf_free(m, v148)
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return int32(0)
																			} else {
																				v154 = v146
																				m.G0 = v8 + int32(32)
																				return v154
																			}
																		}
																	} else {
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																		v143 = F_pushf_flush(m, v142)
																		mBase = m.M
																		v144 = m.ExcPending
																		if v144 != 0 {
																			return int32(0)
																		} else {
																			v146 = v143
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																			if v148 == int32(0) {
																				v154 = v146
																				m.G0 = v8 + int32(32)
																				return v154
																			} else {
																				F_pushf_free(m, v148)
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return int32(0)
																				} else {
																					v154 = v146
																					m.G0 = v8 + int32(32)
																					return v154
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															v126 = F_pgp_mpi_write(m, v103, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = v126
																v129 = F_pgp_mpi_free(m, v116)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return int32(0)
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
																	v132 = F_pgp_mpi_free(m, v131)
																	mBase = m.M
																	v133 = m.ExcPending
																	if v133 != 0 {
																		return int32(0)
																	} else {
																		v135 = v128
																		if v135 < int32(0) {
																			v146 = v135
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																			if v148 == int32(0) {
																				v154 = v146
																				m.G0 = v8 + int32(32)
																				return v154
																			} else {
																				F_pushf_free(m, v148)
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return int32(0)
																				} else {
																					v154 = v146
																					m.G0 = v8 + int32(32)
																					return v154
																				}
																			}
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																			v143 = F_pushf_flush(m, v142)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return int32(0)
																			} else {
																				v146 = v143
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																				if v148 == int32(0) {
																					v154 = v146
																					m.G0 = v8 + int32(32)
																					return v154
																				} else {
																					F_pushf_free(m, v148)
																					mBase = m.M
																					v152 = m.ExcPending
																					if v152 != 0 {
																						return int32(0)
																					} else {
																						v154 = v146
																						m.G0 = v8 + int32(32)
																						return v154
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
										default:
											v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v143 = F_pushf_flush(m, v142)
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												v146 = v143
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
												if v148 == int32(0) {
													v154 = v146
													m.G0 = v8 + int32(32)
													return v154
												} else {
													F_pushf_free(m, v148)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														v154 = v146
														m.G0 = v8 + int32(32)
														return v154
													}
												}
											}
										case 15:
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v61 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v61
											*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v61
											*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v61
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											v73 = F_create_secmsg(m, l0, v8+int32(28), v70-int32(1))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												if v73 < int32(0) {
													v94 = v73
													v95 = F_pgp_mpi_free(m, v75)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
														v98 = F_pgp_mpi_free(m, v97)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
															v101 = F_pgp_mpi_free(m, v100)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int32(0)
															} else {
																v135 = v94
																if v135 < int32(0) {
																	v146 = v135
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																	if v148 == int32(0) {
																		v154 = v146
																		m.G0 = v8 + int32(32)
																		return v154
																	} else {
																		F_pushf_free(m, v148)
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
																			return int32(0)
																		} else {
																			v154 = v146
																			m.G0 = v8 + int32(32)
																			return v154
																		}
																	}
																} else {
																	v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																	v143 = F_pushf_flush(m, v142)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return int32(0)
																	} else {
																		v146 = v143
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																		if v148 == int32(0) {
																			v154 = v146
																			m.G0 = v8 + int32(32)
																			return v154
																		} else {
																			F_pushf_free(m, v148)
																			mBase = m.M
																			v152 = m.ExcPending
																			if v152 != 0 {
																				return int32(0)
																			} else {
																				v154 = v146
																				m.G0 = v8 + int32(32)
																				return v154
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v82 = F_pgp_elgamal_encrypt(m, v10, v75, v8+int32(24), v8+int32(20))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														if v82 < int32(0) {
															v94 = v82
															v95 = F_pgp_mpi_free(m, v75)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
																v98 = F_pgp_mpi_free(m, v97)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
																	v101 = F_pgp_mpi_free(m, v100)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int32(0)
																	} else {
																		v135 = v94
																		if v135 < int32(0) {
																			v146 = v135
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																			if v148 == int32(0) {
																				v154 = v146
																				m.G0 = v8 + int32(32)
																				return v154
																			} else {
																				F_pushf_free(m, v148)
																				mBase = m.M
																				v152 = m.ExcPending
																				if v152 != 0 {
																					return int32(0)
																				} else {
																					v154 = v146
																					m.G0 = v8 + int32(32)
																					return v154
																				}
																			}
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																			v143 = F_pushf_flush(m, v142)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return int32(0)
																			} else {
																				v146 = v143
																				v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																				if v148 == int32(0) {
																					v154 = v146
																					m.G0 = v8 + int32(32)
																					return v154
																				} else {
																					F_pushf_free(m, v148)
																					mBase = m.M
																					v152 = m.ExcPending
																					if v152 != 0 {
																						return int32(0)
																					} else {
																						v154 = v146
																						m.G0 = v8 + int32(32)
																						return v154
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															v87 = F_pgp_mpi_write(m, v60, v86)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																if v87 < int32(0) {
																	v94 = v87
																	v95 = F_pgp_mpi_free(m, v75)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
																		v98 = F_pgp_mpi_free(m, v97)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
																			v101 = F_pgp_mpi_free(m, v100)
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
																				return int32(0)
																			} else {
																				v135 = v94
																				if v135 < int32(0) {
																					v146 = v135
																					v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																					if v148 == int32(0) {
																						v154 = v146
																						m.G0 = v8 + int32(32)
																						return v154
																					} else {
																						F_pushf_free(m, v148)
																						mBase = m.M
																						v152 = m.ExcPending
																						if v152 != 0 {
																							return int32(0)
																						} else {
																							v154 = v146
																							m.G0 = v8 + int32(32)
																							return v154
																						}
																					}
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																					v143 = F_pushf_flush(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						v146 = v143
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																						if v148 == int32(0) {
																							v154 = v146
																							m.G0 = v8 + int32(32)
																							return v154
																						} else {
																							F_pushf_free(m, v148)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return int32(0)
																							} else {
																								v154 = v146
																								m.G0 = v8 + int32(32)
																								return v154
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
																	v92 = F_pgp_mpi_write(m, v60, v91)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		v94 = v92
																		v95 = F_pgp_mpi_free(m, v75)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int32(0)
																		} else {
																			v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
																			v98 = F_pgp_mpi_free(m, v97)
																			mBase = m.M
																			v99 = m.ExcPending
																			if v99 != 0 {
																				return int32(0)
																			} else {
																				v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
																				v101 = F_pgp_mpi_free(m, v100)
																				mBase = m.M
																				v102 = m.ExcPending
																				if v102 != 0 {
																					return int32(0)
																				} else {
																					v135 = v94
																					if v135 < int32(0) {
																						v146 = v135
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																						if v148 == int32(0) {
																							v154 = v146
																							m.G0 = v8 + int32(32)
																							return v154
																						} else {
																							F_pushf_free(m, v148)
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return int32(0)
																							} else {
																								v154 = v146
																								m.G0 = v8 + int32(32)
																								return v154
																							}
																						}
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																						v143 = F_pushf_flush(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							v146 = v143
																							v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																							if v148 == int32(0) {
																								v154 = v146
																								m.G0 = v8 + int32(32)
																								return v154
																							} else {
																								F_pushf_free(m, v148)
																								mBase = m.M
																								v152 = m.ExcPending
																								if v152 != 0 {
																									return int32(0)
																								} else {
																									v154 = v146
																									m.G0 = v8 + int32(32)
																									return v154
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
