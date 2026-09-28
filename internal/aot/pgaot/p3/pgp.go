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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
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
														if v37 != 0 {
															base.MemoryFill(m, v42, int32(0), v37)
														} else {
														}
														F_pfree(m, v42)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															if v40 != 0 {
																if v37 != 0 {
																	base.MemoryFill(m, v40, int32(0), v37)
																} else {
																}
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
																		if v37 != 0 {
																			base.MemoryFill(m, v38, int32(0), v37)
																		} else {
																		}
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
																	if v37 != 0 {
																		base.MemoryFill(m, v38, int32(0), v37)
																	} else {
																	}
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
															if v37 != 0 {
																base.MemoryFill(m, v40, int32(0), v37)
															} else {
															}
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
																	if v37 != 0 {
																		base.MemoryFill(m, v38, int32(0), v37)
																	} else {
																	}
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
																if v37 != 0 {
																	base.MemoryFill(m, v38, int32(0), v37)
																} else {
																}
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
															if v37 != 0 {
																base.MemoryFill(m, v42, int32(0), v37)
															} else {
															}
															F_pfree(m, v42)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																if v40 != 0 {
																	if v37 != 0 {
																		base.MemoryFill(m, v40, int32(0), v37)
																	} else {
																	}
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
																			if v37 != 0 {
																				base.MemoryFill(m, v38, int32(0), v37)
																			} else {
																			}
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
																		if v37 != 0 {
																			base.MemoryFill(m, v38, int32(0), v37)
																		} else {
																		}
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
																if v37 != 0 {
																	base.MemoryFill(m, v40, int32(0), v37)
																} else {
																}
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
																		if v37 != 0 {
																			base.MemoryFill(m, v38, int32(0), v37)
																		} else {
																		}
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
																	if v37 != 0 {
																		base.MemoryFill(m, v38, int32(0), v37)
																	} else {
																	}
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
																if v37 != 0 {
																	base.MemoryFill(m, v42, int32(0), v37)
																} else {
																}
																F_pfree(m, v42)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	if v40 != 0 {
																		if v37 != 0 {
																			base.MemoryFill(m, v40, int32(0), v37)
																		} else {
																		}
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
																				if v37 != 0 {
																					base.MemoryFill(m, v38, int32(0), v37)
																				} else {
																				}
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
																			if v37 != 0 {
																				base.MemoryFill(m, v38, int32(0), v37)
																			} else {
																			}
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
																	if v37 != 0 {
																		base.MemoryFill(m, v40, int32(0), v37)
																	} else {
																	}
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
																			if v37 != 0 {
																				base.MemoryFill(m, v38, int32(0), v37)
																			} else {
																			}
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
																		if v37 != 0 {
																			base.MemoryFill(m, v38, int32(0), v37)
																		} else {
																		}
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
																	if v37 != 0 {
																		base.MemoryFill(m, v42, int32(0), v37)
																	} else {
																	}
																	F_pfree(m, v42)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		if v40 != 0 {
																			if v37 != 0 {
																				base.MemoryFill(m, v40, int32(0), v37)
																			} else {
																			}
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
																					if v37 != 0 {
																						base.MemoryFill(m, v38, int32(0), v37)
																					} else {
																					}
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
																				if v37 != 0 {
																					base.MemoryFill(m, v38, int32(0), v37)
																				} else {
																				}
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
																		if v37 != 0 {
																			base.MemoryFill(m, v40, int32(0), v37)
																		} else {
																		}
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
																				if v37 != 0 {
																					base.MemoryFill(m, v38, int32(0), v37)
																				} else {
																				}
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
																			if v37 != 0 {
																				base.MemoryFill(m, v38, int32(0), v37)
																			} else {
																			}
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
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v147
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
		v147 = v20
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_mpi_check(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v27 == int32(0) {
		v147 = v20
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v31 = F_mpi_check(m, v17)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v31 == int32(0) {
		v147 = v20
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v35 = F_mpi_check(m, v16)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v35 == int32(0) {
		v147 = v20
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v39 <= int32(_a_F_pgp_elgamal_encrypt_0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v63 == int32(0) {
		v147 = v133
		goto L1
	} else {
		goto L63
	}
L12:
	;
	v55 = int32(1)
	v59 = base.I32_div_s(v54+int32(7), int32(8))
	if v59 <= v55 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v43 = base.I32_div_s(v39, int32(10))
	v54 = v43 + int32(160)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v46 = int32(3)
	v54 = int32(base.Ui32(int32(base.Ui32(v39)>>(uint(v46)%32))*v46+int32(600)) >> (uint(int32(1)) % 32))
	goto L12
L16:
	;
	v62 = v55
	goto L18
L17:
	;
	v62 = v59
	goto L18
L18:
	;
	v63 = F_palloc(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v65 = m.Env.Pgmem_bn_rand(m, v54, v63, v59)
	mBase = m.M
	if v65 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v133 = int32(-109)
	goto L11
L21:
	;
	goto L22
L22:
	;
	v69 = int32(1)
	if v19 <= v69 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = v69
	goto L25
L24:
	;
	v72 = v19
	goto L25
L25:
	;
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v75 = F_palloc(m, v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v77 = F_palloc(m, v72)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v79 = int32(-109)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v85 = m.Env.Pgmem_bn_op(m, int32(1), v81, v82, v63, v65, v83, v84, v73, v19)
	mBase = m.M
	if v85 < int32(0) {
		v116 = v79
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v77 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v93 = m.Env.Pgmem_bn_op(m, int32(1), v89, v90, v63, v65, v91, v92, v75, v19)
	mBase = m.M
	if v93 < int32(0) {
		v116 = v79
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v101 = m.Env.Pgmem_bn_op(m, int32(2), v97, v98, v75, v93, v99, v100, v77, v19)
	mBase = m.M
	if v101 < int32(0) {
		v116 = v79
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v104 = F_bytes_to_mpi(m, v73, v85)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
	v107 = F_bytes_to_mpi(m, v77, v101)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v107
	if v107 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v112 = int32(0)
	goto L37
L36:
	;
	v112 = int32(-109)
	goto L37
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v114 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = v112
	goto L40
L39:
	;
	v115 = int32(-109)
	goto L40
L40:
	;
	v116 = v115
	goto L29
L41:
	;
	if v72 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	if v75 != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	F_pfree(m, v77)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	base.MemoryFill(m, v77, int32(0), v72)
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L43
L49:
	;
	if v72 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	if v73 == int32(0) {
		v133 = v116
		goto L11
	} else {
		goto L57
	}
L52:
	;
	F_pfree(m, v75)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L56
	}
L53:
	;
	base.MemoryFill(m, v75, int32(0), v72)
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L51
L57:
	;
	if v72 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_pfree(m, v73)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L62
	}
L59:
	;
	base.MemoryFill(m, v73, int32(0), v72)
	goto L61
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v133 = v116
	goto L11
L63:
	;
	if v62 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_pfree(m, v63)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L68
	}
L65:
	;
	base.MemoryFill(m, v63, int32(0), v62)
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	v147 = v133
	goto L1
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
	var v103 int32
	_ = v103
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
	var v137 int32
	_ = v137
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int64
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
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
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v557 int32
	_ = v557
	var v561 int64
	_ = v561
	var v564 int64
	_ = v564
	var v567 int64
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	v9 = m.G0
	v11 = v9 - int32(288)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(288)
	return v649
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v649 = int32(-13)
	goto L1
L5:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_pushf_free_all(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L6
	} else {
		goto L181
	}
L6:
	;
	return int32(0)
L7:
	;
	if v18 < int32(0) {
		v640 = v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
		v640 = v110
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v103)
	v110 = int32(0)
	goto L16
L18:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+15)))
	v103 = v96&int32(31) | int32(96)
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
	v103 = int32(255)
	goto L17
L30:
	;
	v103 = v67
	goto L17
L31:
	;
	goto L32
L32:
	;
	v83 = v67 | int32(1)
	if base.Ui32(v32) <= base.Ui32((v83&int32(15)|int32(16))<<(uint(v79)%32)) {
		v103 = v83
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
		v640 = v123
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
		v640 = v350
		goto L5
	} else {
		goto L94
	}
L39:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v229 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v229)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v231)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v233)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+19)) = uint8(v235)
	v238 = v11 + int32(16)
	if v233 == int32(0) {
		v253 = v238 | v229
		goto L66
	} else {
		goto L67
	}
L40:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v217
	if v217 == int32(0) {
		goto L39
	} else {
		goto L65
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
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
	v137 = v134 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v137))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v137)%32))&int32(1) == int32(0)) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v154
	v159 = int32(0)
	v163 = m.G0
	v165 = v163 - int32(16)
	m.G0 = v165
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v159
	v171 = F_open(m, int32(_a_F_pgp_encrypt_0), v159, v165)
	mBase = m.M
	if v171 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v154 = int32(0)
	goto L48
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v137<<(uint(int32(2))%32))+uint32(_c_F_pgp_encrypt[0])))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v154 = v153
	goto L48
L48:
	;
	goto L45
L49:
	;
	if v204 == int32(0) {
		v640 = int32(-17)
		goto L5
	} else {
		goto L62
	}
L50:
	;
	v174 = int32(1)
	if v154 == int32(0) {
		v197 = v174
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v204 = v159
	goto L52
L52:
	;
	m.G0 = v165 + int32(16)
	goto L49
L53:
	;
	v199 = F_close(m, v171)
	mBase = m.M
	v204 = v197
	goto L52
L54:
	;
	v177 = l0 + int32(136)
	v178 = v154
	goto L55
L55:
	;
	v183 = F_read(m, v171, v177, v178)
	mBase = m.M
	if v183 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v197 = v174
	goto L53
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_encrypt[1]))
	if v187 == int32(27) {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v192 = v178 - v183
	if v192 != 0 {
		v177 = v177 + v183
		v178 = v192
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v197 = int32(0)
	goto L53
L61:
	;
	goto L56
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v211 == int32(0) {
		goto L39
	} else {
		goto L63
	}
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v215 = F_pgp_write_pubenc_sesskey(m, l0, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v350 = v215
	goto L38
L65:
	;
	base.MemoryCopy(m, l0+int32(136), l0+int32(11), v217)
	goto L39
L66:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v254 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(l0)+2))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+20)) = v243
	if v233 != int32(3) {
		v253 = v238 | int32(12)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v249)
	v253 = v238 | int32(13)
	goto L66
L69:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v255)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	v262 = int32(0)
	v265 = F_pgp_cfb_create(m, v11+int32(276), v231, l0+int32(11), v261, v262, v262, v262)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	v293 = v253
	goto L71
L71:
	;
	v294 = int32(195)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+276)) = uint8(v294)
	v300 = v293 - (v11 + int32(16))
	if v300 <= int32(191) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v265 < int32(0) {
		v350 = v265
		goto L38
	} else {
		goto L73
	}
L73:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v273 = F_pgp_cfb_encrypt(m, v269, v11+int32(8), int32(1), v253)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v281 = F_pgp_cfb_encrypt(m, v275, l0+int32(136), v278, v253+int32(1))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	F_pgp_cfb_free(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v288 = v286 + int32(1)
	if v288 < int32(0) {
		v350 = v288
		goto L38
	} else {
		goto L77
	}
L77:
	;
	v293 = v288 + v253
	goto L71
L78:
	;
	v337 = F_pushf_write(m, v228, v11+int32(276), v333-(v11+int32(276)))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L85
	}
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v300)
	v333 = v11 + int32(278)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(v300) <= base.Ui32(int32(_a_F_pgp_encrypt_1)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v309 = v300 - int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+278)) = uint8(v309)
	v314 = int32(base.Ui32(v309)>>(uint(int32(8))%32)) + int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v314)
	v333 = v11 + int32(279)
	goto L78
L83:
	;
	goto L84
L84:
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
	goto L78
L85:
	;
	if int32(0) <= v337 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v343 = F_pushf_write(m, v228, v11+int32(16), v300)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	v345 = v337
	goto L88
L88:
	;
	if v300 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v345 = v343
	goto L88
L90:
	;
	v350 = v345
	goto L38
L91:
	;
	base.MemoryFill(m, v11+int32(16), int32(0), v300)
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L90
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
		v640 = v365
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v370 = v11 + int32(8)
	v372 = F_pushf_create(m, v370, int32(_a_F_pgp_encrypt_2), l0, v356)
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
		v640 = v372
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v376
	v379 = F_pushf_create(m, v370, int32(_a_F_pgp_encrypt_3), l0, v376)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	if v379 < int32(0) {
		v640 = v379
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v385 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v389 = F_pushf_create(m, v370, int32(_a_F_pgp_encrypt_4), l0, v383)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L107
	}
L105:
	;
	v395 = v383
	goto L106
L106:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v401 = v398 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v401))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v401)%32))&int32(1) == int32(0)) != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v389 < int32(0) {
		v640 = v389
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v393
	v395 = v393
	goto L106
L109:
	;
	v419 = int32(0)
	v423 = m.G0
	v425 = v423 - int32(16)
	m.G0 = v425
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v419
	v431 = F_open(m, int32(_a_F_pgp_encrypt_0), v419, v425)
	mBase = m.M
	if v431 != int32(-1) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v418 = int32(0)
	goto L112
L111:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v401<<(uint(int32(2))%32))+uint32(_c_F_pgp_encrypt[0])))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	v418 = v417
	goto L112
L112:
	;
	goto L109
L113:
	;
	if v464 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v434 = int32(1)
	if v418 == int32(0) {
		v457 = v434
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v464 = v419
	goto L116
L116:
	;
	m.G0 = v425 + int32(16)
	goto L113
L117:
	;
	v459 = F_close(m, v431)
	mBase = m.M
	v464 = v457
	goto L116
L118:
	;
	v437 = v11 + int32(16)
	v438 = v418
	goto L119
L119:
	;
	v443 = F_read(m, v431, v437, v438)
	mBase = m.M
	if v443 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v457 = v434
	goto L117
L121:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_encrypt[1]))
	if v447 == int32(27) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v452 = v438 - v443
	if v452 != 0 {
		v437 = v437 + v443
		v438 = v452
		goto L119
	} else {
		goto L125
	}
L124:
	;
	v457 = int32(0)
	goto L117
L125:
	;
	goto L120
L126:
	;
	v640 = int32(-17)
	goto L5
L127:
	;
	goto L128
L128:
	;
	v473 = v11 + int32(16)
	v474 = v473 + v418
	v475 = int32(2)
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474-v475))))
	*(*uint16)(unsafe.Add(mBase, uint32(v474))) = uint16(v477)
	v480 = v418 + v475
	v481 = F_pushf_write(m, v395, v473, v480)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	if v480 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v481 < int32(0) {
		v640 = v481
		goto L5
	} else {
		goto L134
	}
L131:
	;
	base.MemoryFill(m, v473, int32(0), v480)
	goto L133
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v487 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v548 != 0 {
		goto L153
	} else {
		goto L154
	}
L136:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v545 = v490
	goto L135
L137:
	;
	goto L138
L138:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v492 <= int32(0) {
		v545 = v491
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v497 = m.G0
	v499 = v497 - int32(16)
	m.G0 = v499
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v499)+14)) = uint8(v501)
	v503 = int32(200)
	*(*uint8)(unsafe.Add(mBase, uint32(v499)+15)) = uint8(v503)
	v508 = F_pushf_write(m, v491, v499+int32(15), int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L6
	} else {
		goto L141
	}
L140:
	;
	m.G0 = v499 + int32(16)
	if v536 < int32(0) {
		v640 = v536
		goto L5
	} else {
		goto L152
	}
L141:
	;
	if v508 < int32(0) {
		v536 = v508
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v515 = F_pushf_create(m, v499+int32(8), int32(_a_F_pgp_encrypt_2), l0, v491)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v515 < int32(0) {
		v536 = v515
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	v523 = F_pushf_write(m, v519, v499+int32(14), int32(1))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	if int32(0) <= v523 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	v528 = F_pgp_compress_filter(m, v11+int32(8), l0, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	v532 = v523
	goto L148
L148:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v499)+8))
	F_pushf_free(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L151
	}
L149:
	;
	if int32(0) <= v528 {
		v536 = v528
		goto L140
	} else {
		goto L150
	}
L150:
	;
	v532 = v528
	goto L148
L151:
	;
	v536 = v532
	goto L140
L152:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v542
	v545 = v542
	goto L135
L153:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v551 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v554 = int32(98)
	goto L155
L155:
	;
	v555 = F_time(m)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v555)
	v557 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v554)
	v561 = int64(base.Ui64(v555) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v561)
	v564 = int64(base.Ui64(v555) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+19)) = uint8(v564)
	v567 = int64(base.Ui64(v555) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v567)
	v569 = int32(203)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+287)) = uint8(v569)
	v574 = F_pushf_write(m, v545, v11+int32(287), int32(1))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L6
	} else {
		goto L159
	}
L156:
	;
	v552 = int32(117)
	goto L158
L157:
	;
	v552 = int32(116)
	goto L158
L158:
	;
	v554 = v552
	goto L155
L159:
	;
	if v574 < int32(0) {
		v640 = v574
		goto L5
	} else {
		goto L160
	}
L160:
	;
	v581 = F_pushf_create(m, v11+int32(276), int32(_a_F_pgp_encrypt_2), l0, v545)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	if v581 < int32(0) {
		v640 = v581
		goto L5
	} else {
		goto L162
	}
L162:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v589 = F_pushf_write(m, v585, v11+int32(16), int32(6))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	if v589 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_pushf_free(m, v591)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v591
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v598 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v640 = v589
	goto L5
L168:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v616 = v614 - v615
	goto L173
L169:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v601 == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v607 = F_pushf_create(m, v11+int32(8), int32(_a_F_pgp_encrypt_5), l0, v591)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v607 < int32(0) {
		v640 = v607
		goto L5
	} else {
		goto L172
	}
L172:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v611
	goto L168
L173:
	;
	v620 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v620)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v623
	v625 = v622 - v623
	if v616 < v625 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v632 = F_pushf_write(m, v630, v631, v627)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L178
	}
L175:
	;
	v627 = v616
	goto L177
L176:
	;
	v627 = v625
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v623 + v627
	goto L174
L178:
	;
	if v632 < int32(0) {
		v640 = v632
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v637 = F_pushf_flush(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	v640 = v637
	goto L5
L181:
	;
	v649 = v640
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
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
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
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
	return v308
L2:
	;
	return int32(0)
L3:
	;
	if v19 < int32(0) {
		v308 = v19
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = v3
	v31 = v3
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
		v155 = v30
		v156 = v31
		v158 = v33
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
		v155 = v30
		v156 = v31
		v158 = v33
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
		v127 = v33
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
		v137 = v30
		v138 = v31
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
	v137 = v30
	v138 = v31
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
	v127 = v33 + int32(1)
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
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v134 = int32(-114)
	v135 = v58
	v137 = v30
	v138 = v31
	v139 = int32(1)
	v140 = v33
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
	v137 = v30
	v138 = v31
	v139 = int32(1)
	v140 = v33
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
	if int32(1)<<(uint(v81)%32)&int32(_a_F_pgp_get_keyid_0) != 0 {
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
	v137 = v30
	v138 = v31
	v139 = v32
	v140 = v33
	goto L12
L34:
	;
	v98 = int32(1)
	v134 = v98
	v135 = int32(0)
	v137 = v30 + v98
	v138 = v31
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
	v114 = v31 + int32(1)
	if v107 < int32(0) {
		v134 = v112
		v135 = v102
		v137 = v30
		v138 = v114
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
		v137 = v30
		v138 = v114
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
		v137 = v30
		v138 = v114
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
	v137 = v30
	v138 = v114
	v139 = v32
	v140 = v33
	goto L12
L44:
	;
	v134 = v130
	v135 = int32(0)
	v137 = v30
	v138 = v31
	v139 = v32
	v140 = v127
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
		v30 = v137
		v31 = v138
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
	v155 = v137
	v156 = v138
	v158 = v140
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
		v308 = v152
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
	if v156 != 0 {
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
	if v155 != 0 {
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
	if int32(1) < v155 {
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
	if int32(1) < v156 {
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
		v308 = v178
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v155|v156 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v308 = int32(6)
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
	if v158 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_get_keyid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+3)) = v186
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_get_keyid[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v189
	goto L69
L74:
	;
	goto L75
L75:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v192 = int32(15)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v196)
	v198 = int32(4)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v191)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v202)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)) = uint8(v209)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v204)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v215)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v222)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v217)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v228)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)) = uint8(v235)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v230)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v241)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v248)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v243)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v254)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v261)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v256)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v267)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v274)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v269)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v280)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)))
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v283)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282&v192)+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)) = uint8(v289)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v282)>>(uint(v198)%32)))+uint32(_c_F_pgp_get_keyid[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)) = uint8(v295)
	v308 = int32(16)
	goto L1
L76:
	;
	v308 = int32(-119)
	goto L1
L77:
	;
	goto L78
L78:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_get_keyid[3]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+3)) = v302
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_get_keyid[4]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v305
	goto L69
}
func F_pgp_load_digest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	v5 = int32(1)
	v6 = l0 - v5
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v6))|base.B2i32(int32(base.Ui32(int32(903))>>(uint(v6)%32))&v5 == int32(0)) != 0 {
		v29 = int32(-100)
		return v29
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_c_F_pgp_load_digest[0])))
		v24 = F_px_find_digest(m, v23, l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 != 0 {
				v28 = int32(-104)
			} else {
				v28 = int32(0)
			}
			v29 = v28
			return v29
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
	var v39 int32
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if base.Ui32(int32(_a_F_pgp_mpi_create_0)) <= base.Ui32(l1) {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
		F_px_debug(m, int32(_a_F_pgp_mpi_create_1), v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v39 = int32(-100)
			m.G0 = v9 + int32(16)
			return v39
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
				base.MemoryCopy(m, v31, l0, v23)
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
			v39 = int32(0)
			m.G0 = v9 + int32(16)
			return v39
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
	var v45 int32
	_ = v45
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
						v45 = v43 + int32(12)
						if v45 != 0 {
							base.MemoryFill(m, v31, int32(0), v45)
						} else {
						}
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
func F_pgp_pub_decrypt_bytea(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(v33)
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
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
func F_pgp_pub_decrypt_text(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
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
	var v94 int32
	_ = v94
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v17 < int32(3) {
				v29 = v2
				v30 = v2
				v31 = int32(1)
				v33 = F_decrypt_internal(m, v31, v31, v10, v15, v29, v30)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = int32(1)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
					v39 = v37 & v35
					if v39 != 0 {
						v40 = v35
					} else {
						v40 = int32(4)
					}
					if v37 == int32(1) {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
						if v47 == int32(18) {
							v50 = int32(16)
						} else {
							v50 = int32(0)
						}
						if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v57 = int32(4)
						} else {
							v57 = v50
						}
						v68 = v57
					} else {
						v58 = int32(1)
						if v39 != 0 {
							v68 = int32(base.Ui32(v37)>>(uint(v58)%32)) - v58
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v68 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v33+v40, v68)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v71 != v10 {
							F_pfree(m, v10)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v75 != v15 {
									F_pfree(m, v15)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int64(0)
									} else {
										v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v79 < int32(3) {
											return base.I64_extend_i32_u(v33)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v82 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int64(0)
												} else {
													v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v87 = v86
													if base.I32_extend16_s(v87) < int32(4) {
														return base.I64_extend_i32_u(v33)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														if v30 == v91 {
															return base.I64_extend_i32_u(v33)
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v33)
															}
														}
													}
												}
											} else {
												v87 = v79
												if base.I32_extend16_s(v87) < int32(4) {
													return base.I64_extend_i32_u(v33)
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													if v30 == v91 {
														return base.I64_extend_i32_u(v33)
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v33)
														}
													}
												}
											}
										}
									}
								} else {
									v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v79 < int32(3) {
										return base.I64_extend_i32_u(v33)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v82 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int64(0)
											} else {
												v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v87 = v86
												if base.I32_extend16_s(v87) < int32(4) {
													return base.I64_extend_i32_u(v33)
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													if v30 == v91 {
														return base.I64_extend_i32_u(v33)
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v33)
														}
													}
												}
											}
										} else {
											v87 = v79
											if base.I32_extend16_s(v87) < int32(4) {
												return base.I64_extend_i32_u(v33)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
												if v30 == v91 {
													return base.I64_extend_i32_u(v33)
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v33)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v75 != v15 {
								F_pfree(m, v15)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v79 < int32(3) {
										return base.I64_extend_i32_u(v33)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v82 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int64(0)
											} else {
												v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v87 = v86
												if base.I32_extend16_s(v87) < int32(4) {
													return base.I64_extend_i32_u(v33)
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													if v30 == v91 {
														return base.I64_extend_i32_u(v33)
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v33)
														}
													}
												}
											}
										} else {
											v87 = v79
											if base.I32_extend16_s(v87) < int32(4) {
												return base.I64_extend_i32_u(v33)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
												if v30 == v91 {
													return base.I64_extend_i32_u(v33)
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v33)
													}
												}
											}
										}
									}
								}
							} else {
								v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v79 < int32(3) {
									return base.I64_extend_i32_u(v33)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v82 != v29 {
										F_pfree(m, v29)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int64(0)
										} else {
											v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
											v87 = v86
											if base.I32_extend16_s(v87) < int32(4) {
												return base.I64_extend_i32_u(v33)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
												if v30 == v91 {
													return base.I64_extend_i32_u(v33)
												} else {
													F_pfree(m, v30)
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v33)
													}
												}
											}
										}
									} else {
										v87 = v79
										if base.I32_extend16_s(v87) < int32(4) {
											return base.I64_extend_i32_u(v33)
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
											if v30 == v91 {
												return base.I64_extend_i32_u(v33)
											} else {
												F_pfree(m, v30)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v33)
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
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v21 = F_pg_detoast_datum_packed(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
					if v23 < int32(4) {
						v29 = v21
						v30 = v2
						v31 = int32(1)
						v33 = F_decrypt_internal(m, v31, v31, v10, v15, v29, v30)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = int32(1)
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
							v39 = v37 & v35
							if v39 != 0 {
								v40 = v35
							} else {
								v40 = int32(4)
							}
							if v37 == int32(1) {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
								if v47 == int32(18) {
									v50 = int32(16)
								} else {
									v50 = int32(0)
								}
								if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v57 = int32(4)
								} else {
									v57 = v50
								}
								v68 = v57
							} else {
								v58 = int32(1)
								if v39 != 0 {
									v68 = int32(base.Ui32(v37)>>(uint(v58)%32)) - v58
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v68 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_pg_verifymbstr(m, v33+v40, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v71 != v10 {
									F_pfree(m, v10)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int64(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v75 != v15 {
											F_pfree(m, v15)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v79 < int32(3) {
													return base.I64_extend_i32_u(v33)
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v82 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int64(0)
														} else {
															v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v87 = v86
															if base.I32_extend16_s(v87) < int32(4) {
																return base.I64_extend_i32_u(v33)
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
																if v30 == v91 {
																	return base.I64_extend_i32_u(v33)
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int64(0)
																	} else {
																		return base.I64_extend_i32_u(v33)
																	}
																}
															}
														}
													} else {
														v87 = v79
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												}
											}
										} else {
											v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v79 < int32(3) {
												return base.I64_extend_i32_u(v33)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v82 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int64(0)
													} else {
														v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v87 = v86
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												} else {
													v87 = v79
													if base.I32_extend16_s(v87) < int32(4) {
														return base.I64_extend_i32_u(v33)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														if v30 == v91 {
															return base.I64_extend_i32_u(v33)
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v33)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v75 != v15 {
										F_pfree(m, v15)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int64(0)
										} else {
											v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v79 < int32(3) {
												return base.I64_extend_i32_u(v33)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v82 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int64(0)
													} else {
														v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v87 = v86
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												} else {
													v87 = v79
													if base.I32_extend16_s(v87) < int32(4) {
														return base.I64_extend_i32_u(v33)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														if v30 == v91 {
															return base.I64_extend_i32_u(v33)
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v33)
															}
														}
													}
												}
											}
										}
									} else {
										v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v79 < int32(3) {
											return base.I64_extend_i32_u(v33)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v82 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int64(0)
												} else {
													v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v87 = v86
													if base.I32_extend16_s(v87) < int32(4) {
														return base.I64_extend_i32_u(v33)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														if v30 == v91 {
															return base.I64_extend_i32_u(v33)
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v33)
															}
														}
													}
												}
											} else {
												v87 = v79
												if base.I32_extend16_s(v87) < int32(4) {
													return base.I64_extend_i32_u(v33)
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
													if v30 == v91 {
														return base.I64_extend_i32_u(v33)
													} else {
														F_pfree(m, v30)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v33)
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
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v27 = F_pg_detoast_datum_packed(m, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							v29 = v21
							v30 = v27
							v31 = int32(1)
							v33 = F_decrypt_internal(m, v31, v31, v10, v15, v29, v30)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								v35 = int32(1)
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								v39 = v37 & v35
								if v39 != 0 {
									v40 = v35
								} else {
									v40 = int32(4)
								}
								if v37 == int32(1) {
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
									if v47 == int32(18) {
										v50 = int32(16)
									} else {
										v50 = int32(0)
									}
									if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v57 = int32(4)
									} else {
										v57 = v50
									}
									v68 = v57
								} else {
									v58 = int32(1)
									if v39 != 0 {
										v68 = int32(base.Ui32(v37)>>(uint(v58)%32)) - v58
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v68 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_pg_verifymbstr(m, v33+v40, v68)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v71 != v10 {
										F_pfree(m, v10)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int64(0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v75 != v15 {
												F_pfree(m, v15)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v79 < int32(3) {
														return base.I64_extend_i32_u(v33)
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v82 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int64(0)
															} else {
																v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
																v87 = v86
																if base.I32_extend16_s(v87) < int32(4) {
																	return base.I64_extend_i32_u(v33)
																} else {
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
																	if v30 == v91 {
																		return base.I64_extend_i32_u(v33)
																	} else {
																		F_pfree(m, v30)
																		mBase = m.M
																		v94 = m.ExcPending
																		if v94 != 0 {
																			return int64(0)
																		} else {
																			return base.I64_extend_i32_u(v33)
																		}
																	}
																}
															}
														} else {
															v87 = v79
															if base.I32_extend16_s(v87) < int32(4) {
																return base.I64_extend_i32_u(v33)
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
																if v30 == v91 {
																	return base.I64_extend_i32_u(v33)
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int64(0)
																	} else {
																		return base.I64_extend_i32_u(v33)
																	}
																}
															}
														}
													}
												}
											} else {
												v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v79 < int32(3) {
													return base.I64_extend_i32_u(v33)
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v82 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int64(0)
														} else {
															v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v87 = v86
															if base.I32_extend16_s(v87) < int32(4) {
																return base.I64_extend_i32_u(v33)
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
																if v30 == v91 {
																	return base.I64_extend_i32_u(v33)
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int64(0)
																	} else {
																		return base.I64_extend_i32_u(v33)
																	}
																}
															}
														}
													} else {
														v87 = v79
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												}
											}
										}
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v75 != v15 {
											F_pfree(m, v15)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v79 < int32(3) {
													return base.I64_extend_i32_u(v33)
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v82 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int64(0)
														} else {
															v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v87 = v86
															if base.I32_extend16_s(v87) < int32(4) {
																return base.I64_extend_i32_u(v33)
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
																if v30 == v91 {
																	return base.I64_extend_i32_u(v33)
																} else {
																	F_pfree(m, v30)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int64(0)
																	} else {
																		return base.I64_extend_i32_u(v33)
																	}
																}
															}
														}
													} else {
														v87 = v79
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												}
											}
										} else {
											v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v79 < int32(3) {
												return base.I64_extend_i32_u(v33)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v82 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int64(0)
													} else {
														v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v87 = v86
														if base.I32_extend16_s(v87) < int32(4) {
															return base.I64_extend_i32_u(v33)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
															if v30 == v91 {
																return base.I64_extend_i32_u(v33)
															} else {
																F_pfree(m, v30)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int64(0)
																} else {
																	return base.I64_extend_i32_u(v33)
																}
															}
														}
													}
												} else {
													v87 = v79
													if base.I32_extend16_s(v87) < int32(4) {
														return base.I64_extend_i32_u(v33)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
														if v30 == v91 {
															return base.I64_extend_i32_u(v33)
														} else {
															F_pfree(m, v30)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v33)
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
func F_pgp_pub_encrypt_bytea(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14352(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_pgp_pub_encrypt_text(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn14352(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
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
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
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
	return v799
L2:
	;
	return int32(0)
L3:
	;
	if v24 < int32(0) {
		v799 = v24
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
	v48 = v6
	v49 = v6
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
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	if v779 != 0 {
		goto L206
	} else {
		goto L207
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
	v769 = v59
	v772 = v48
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
	v769 = v67
	v772 = v48
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
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	F_pullf_free(m, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L2
	} else {
		goto L187
	}
L17:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
	switch v216 - int32(1) {
	case 0, 1, 2:
		goto L72
	default:
		goto L70
	case 15, 16:
		v238 = int32(24)
		goto L71
	}
L18:
	;
	v722 = v203
	v725 = int32(0)
	v728 = v49
	v729 = v51
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v71
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_0), v20)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L68
	}
L20:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v196 = F_pgp_skip_packet(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
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
	v722 = int32(-114)
	v725 = v74
	v728 = v49
	v729 = int32(1)
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
	v722 = v78
	v725 = v74
	v728 = v49
	v729 = int32(1)
	goto L16
L28:
	;
	v203 = int32(-116)
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
	v722 = v85
	v725 = v87
	v728 = v87
	v729 = v51
	goto L16
L32:
	;
	v203 = int32(-115)
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
		v722 = v99
		v725 = v92
		v728 = v49
		v729 = v51
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
		v722 = v106
		v725 = v92
		v728 = v49
		v729 = v51
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
		v211 = v91
		goto L17
	} else {
		goto L65
	}
L42:
	;
	v722 = int32(-120)
	v725 = v92
	v728 = v49
	v729 = v51
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
		v203 = v119
		goto L18
	} else {
		goto L46
	}
L46:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+112)))
	v125 = v20 + int32(60)
	v126 = F_pgp_s2k_read(m, v91, v125)
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
		v203 = v126
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v130 = F_pgp_s2k_process(m, v125, v123, l2, l3)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v130 < int32(0) {
		v203 = v130
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v136 = v123 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v136))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v136)%32))&int32(1) == int32(0)) != 0 {
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
	v153 = int32(0)
	goto L54
L53:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v136<<(uint(int32(2))%32))+uint32(_c_F_pgp_set_pubkey[0])))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
	v153 = v152
	goto L54
L54:
	;
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v123
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_1), v20+int32(16))
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
	v164 = v20 + int32(112)
	v165 = F_pullf_read_fixed(m, v91, v153, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L59
	}
L58:
	;
	v203 = int32(-103)
	goto L18
L59:
	;
	if v165 < int32(0) {
		v203 = v165
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+103)))
	v172 = int32(0)
	v174 = F_pgp_cfb_create(m, v20+int32(104), v123, v20+int32(71), v171, v172, v164, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v174 < int32(0) {
		v203 = v174
		goto L18
	} else {
		goto L62
	}
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v182 = F_pullf_create(m, v20+int32(108), int32(_a_F_pgp_set_pubkey_2), v181, v91)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	if v182 < int32(0) {
		v203 = v182
		goto L18
	} else {
		goto L64
	}
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v211 = v186
	goto L17
L65:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_3), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v722 = int32(-118)
	v725 = v92
	v728 = v49
	v729 = v51
	goto L16
L67:
	;
	v722 = v196
	v725 = int32(0)
	v728 = v49
	v729 = v51
	goto L16
L68:
	;
	v203 = int32(-107)
	goto L18
L69:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v711 != 0 {
		goto L175
	} else {
		goto L176
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v216
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_4), v20+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L2
	} else {
		goto L174
	}
L71:
	;
	v240 = F_pgp_mpi_read(m, v211, v238+v215)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L79
	}
L72:
	;
	v221 = F_pgp_mpi_read(m, v211, v215+int32(24))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v221 < int32(0) {
		v707 = v221
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v227 = F_pgp_mpi_read(m, v211, v215+int32(28))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if v227 < int32(0) {
		v707 = v227
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v233 = F_pgp_mpi_read(m, v211, v215+int32(32))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v233 < int32(0) {
		v707 = v233
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v238 = int32(36)
	goto L71
L79:
	;
	if v240 < int32(0) {
		v707 = v240
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
	if v692 < int32(0) {
		v707 = v692
		goto L69
	} else {
		goto L172
	}
L82:
	;
	v249 = F_pullf_read_fixed(m, v211, int32(20), v20+int32(672))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v328 = F_pullf_read_fixed(m, v211, int32(2), v20+int32(672))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L109
	}
L85:
	;
	if v249 < int32(0) {
		v692 = v249
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v256 = F_pgp_load_digest(m, int32(2), v20+int32(636))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L88
	}
L87:
	;
	goto L102
L88:
	;
	if v256 < int32(0) {
		v313 = v256
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
	switch v261 - int32(1) {
	case 0, 1, 2:
		goto L92
	default:
		goto L90
	case 15, 16:
		v277 = int32(24)
		goto L91
	}
L90:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	m.T0[v287].(func(*base.Module, int32, int32))(m, v284, v20+int32(640))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L97
	}
L91:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v215+v277)))
	v281 = F_pgp_mpi_hash(m, v278, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L96
	}
L92:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v266 = F_pgp_mpi_hash(m, v264, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	v270 = F_pgp_mpi_hash(m, v268, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	v274 = F_pgp_mpi_hash(m, v272, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v277 = int32(36)
	goto L91
L96:
	;
	goto L90
L97:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v20)+636))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	m.T0[v291].(func(*base.Module, int32))(m, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v20)+656))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v20)+688))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v20)+640))
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v20)+672))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v20)+648))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v20)+680))
	if base.I64_extend_i32_u(v294^v295)|(v298^v299|(v301^v302)) == int64(0) {
		v313 = v256
		goto L87
	} else {
		goto L99
	}
L99:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_5), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	v313 = int32(-118)
	goto L87
L101:
	;
	goto L106
L102:
	;
	base.MemoryFill(m, v20+int32(672), int32(0), int32(20))
	goto L104
L104:
	;
	goto L101
L105:
	;
	v692 = v313
	goto L81
L106:
	;
	base.MemoryFill(m, v20+int32(640), int32(0), int32(20))
	goto L108
L108:
	;
	goto L105
L109:
	;
	if v328 < int32(0) {
		v692 = v328
		goto L81
	} else {
		goto L110
	}
L110:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+673)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+672)))
	v338 = int32(0)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
	switch v340 - int32(1) {
	case 0, 1, 2:
		goto L113
	default:
		v684 = v338
		goto L111
	case 15, 16:
		v597 = v338
		v598 = int32(24)
		goto L112
	}
L111:
	;
	if v684 == v332|v333<<(uint(int32(8))%32) {
		v692 = v338
		goto L81
	} else {
		goto L170
	}
L112:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v215+v598)))
	v601 = int32(0)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	v613 = v597 + v607>>(uint(int32(8))%32) + v607&int32(255)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	if v614 <= v601 {
		v674 = v613
		goto L157
	} else {
		goto L158
	}
L113:
	;
	v344 = int32(0)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v358 = v344 + v352>>(uint(int32(8))%32) + v352&int32(255)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	if v359 <= v344 {
		v419 = v358
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	v430 = int32(0)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	v442 = v419&int32(_a_F_pgp_set_pubkey_6) + v436>>(uint(int32(8))%32) + v436&int32(255)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	if v443 <= v430 {
		v503 = v442
		goto L129
	} else {
		goto L130
	}
L115:
	;
	goto L114
L116:
	;
	v363 = v359 & int32(3)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if base.Ui32(v359) < base.Ui32(int32(4)) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v403 = v395
	v404 = v396
	v410 = v344
	goto L125
L118:
	;
	v395 = v358
	v396 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v371 = v358
	v372 = int32(0)
	v377 = v344
	goto L121
L121:
	;
	v379 = v372 + v364
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+2)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+3)))
	v387 = v371 + v380 + v382 + v384 + v386
	v388 = int32(4)
	v389 = v372 + v388
	v391 = v377 + v388
	if v391 != v359&int32(2147483644) {
		v371 = v387
		v372 = v389
		v377 = v391
		goto L121
	} else {
		goto L123
	}
L122:
	;
	if v363 == int32(0) {
		v419 = v387
		goto L115
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v395 = v387
	v396 = v389
	goto L117
L125:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v364))))
	v413 = v403 + v412
	v414 = int32(1)
	v417 = v410 + v414
	if v417 != v363 {
		v403 = v413
		v404 = v404 + v414
		v410 = v417
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v419 = v413
	goto L115
L127:
	;
	goto L126
L128:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	v514 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v526 = v503&int32(_a_F_pgp_set_pubkey_6) + v520>>(uint(int32(8))%32) + v520&int32(255)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	if v527 <= v514 {
		v587 = v526
		goto L143
	} else {
		goto L144
	}
L129:
	;
	goto L128
L130:
	;
	v447 = v443 & int32(3)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if base.Ui32(v443) < base.Ui32(int32(4)) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v487 = v479
	v488 = v480
	v494 = v430
	goto L139
L132:
	;
	v479 = v442
	v480 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v455 = v442
	v456 = int32(0)
	v461 = v430
	goto L135
L135:
	;
	v463 = v456 + v448
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+2)))
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+3)))
	v471 = v455 + v464 + v466 + v468 + v470
	v472 = int32(4)
	v473 = v456 + v472
	v475 = v461 + v472
	if v475 != v443&int32(2147483644) {
		v455 = v471
		v456 = v473
		v461 = v475
		goto L135
	} else {
		goto L137
	}
L136:
	;
	if v447 == int32(0) {
		v503 = v471
		goto L129
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v479 = v471
	v480 = v473
	goto L131
L139:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+v448))))
	v497 = v487 + v496
	v498 = int32(1)
	v501 = v494 + v498
	if v501 != v447 {
		v487 = v497
		v488 = v488 + v498
		v494 = v501
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v503 = v497
	goto L129
L141:
	;
	goto L140
L142:
	;
	v597 = v587 & int32(_a_F_pgp_set_pubkey_6)
	v598 = int32(36)
	goto L112
L143:
	;
	goto L142
L144:
	;
	v531 = v527 & int32(3)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	if base.Ui32(v527) < base.Ui32(int32(4)) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v571 = v563
	v572 = v564
	v578 = v514
	goto L153
L146:
	;
	v563 = v526
	v564 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v539 = v526
	v540 = int32(0)
	v545 = v514
	goto L149
L149:
	;
	v547 = v540 + v532
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+1)))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+2)))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+3)))
	v555 = v539 + v548 + v550 + v552 + v554
	v556 = int32(4)
	v557 = v540 + v556
	v559 = v545 + v556
	if v559 != v527&int32(2147483644) {
		v539 = v555
		v540 = v557
		v545 = v559
		goto L149
	} else {
		goto L151
	}
L150:
	;
	if v531 == int32(0) {
		v587 = v555
		goto L143
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v563 = v555
	v564 = v557
	goto L145
L153:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572+v532))))
	v581 = v571 + v580
	v582 = int32(1)
	v585 = v578 + v582
	if v585 != v531 {
		v571 = v581
		v572 = v572 + v582
		v578 = v585
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v587 = v581
	goto L143
L155:
	;
	goto L154
L156:
	;
	v684 = v674 & int32(_a_F_pgp_set_pubkey_6)
	goto L111
L157:
	;
	goto L156
L158:
	;
	v618 = v614 & int32(3)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	if base.Ui32(v614) < base.Ui32(int32(4)) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v658 = v650
	v659 = v651
	v665 = v601
	goto L167
L160:
	;
	v650 = v613
	v651 = int32(0)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v626 = v613
	v627 = int32(0)
	v632 = v601
	goto L163
L163:
	;
	v634 = v627 + v619
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+1)))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+2)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+3)))
	v642 = v626 + v635 + v637 + v639 + v641
	v643 = int32(4)
	v644 = v627 + v643
	v646 = v632 + v643
	if v646 != v614&int32(2147483644) {
		v626 = v642
		v627 = v644
		v632 = v646
		goto L163
	} else {
		goto L165
	}
L164:
	;
	if v618 == int32(0) {
		v674 = v642
		goto L157
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	v650 = v642
	v651 = v644
	goto L159
L167:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659+v619))))
	v668 = v658 + v667
	v669 = int32(1)
	v672 = v665 + v669
	if v672 != v618 {
		v658 = v668
		v659 = v659 + v669
		v665 = v672
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v674 = v668
	goto L157
L169:
	;
	goto L168
L170:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_7), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v692 = int32(-118)
	goto L81
L172:
	;
	v698 = F_pgp_expect_packet_end(m, v211)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	v707 = v698
	goto L69
L174:
	;
	v707 = int32(-118)
	goto L69
L175:
	;
	F_pullf_free(m, v711)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L2
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v714 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	F_pgp_cfb_free(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L2
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v707 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	F_pgp_key_free(m, v215)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L2
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v215
	v722 = v707
	v725 = v215
	v728 = v215
	v729 = v51
	goto L16
L186:
	;
	v722 = v707
	v725 = v92
	v728 = v49
	v729 = v51
	goto L16
L187:
	;
	v734 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v734
	if v725 == v734 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if int32(0) <= v762 {
		v48 = v765
		v49 = v766
		v51 = v729
		goto L5
	} else {
		goto L205
	}
L189:
	;
	v762 = v722
	v764 = v48
	v765 = v48
	v766 = v728
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v722 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v759 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v759
	v762 = v755
	v764 = v757
	v765 = v757
	v766 = v759
	goto L188
L193:
	;
	v755 = v753
	v757 = v48
	goto L192
L194:
	;
	F_pgp_key_free(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L2
	} else {
		goto L204
	}
L195:
	;
	v749 = v722
	v750 = v725
	goto L194
L196:
	;
	goto L197
L197:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v725)+48))
	if v740 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v728 == int32(0) {
		v753 = v746
		goto L193
	} else {
		goto L203
	}
L199:
	;
	v746 = v722
	goto L198
L200:
	;
	goto L201
L201:
	;
	if v48 == int32(0) {
		v755 = v722
		v757 = v725
		goto L192
	} else {
		goto L202
	}
L202:
	;
	v746 = int32(-123)
	goto L198
L203:
	;
	v749 = v746
	v750 = v728
	goto L194
L204:
	;
	v753 = v749
	goto L193
L205:
	;
	v769 = v762
	v772 = v764
	goto L7
L206:
	;
	F_pullf_free(m, v779)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L2
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	if v769 < int32(0) {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	goto L208
L210:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L2
	} else {
		goto L219
	}
L211:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L2
	} else {
		goto L218
	}
L212:
	;
	if v772 == int32(0) {
		v789 = v769
		goto L211
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if v772 != 0 {
		goto L210
	} else {
		goto L217
	}
L215:
	;
	F_pgp_key_free(m, v772)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	v789 = v769
	goto L211
L217:
	;
	v789 = int32(-119)
	goto L211
L218:
	;
	v799 = v789
	goto L1
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v772
	v799 = int32(0)
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
func F_pgp_sym_decrypt_text(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v19 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v23 = F_pg_detoast_datum_packed(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = v23
					v26 = F_decrypt_internal(m, v2, int32(1), v11, v16, int32(0), v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = int32(1)
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
						v32 = v30 & v28
						if v32 != 0 {
							v33 = v28
						} else {
							v33 = int32(4)
						}
						if v30 == int32(1) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
							if v40 == int32(18) {
								v43 = int32(16)
							} else {
								v43 = int32(0)
							}
							if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v50 = int32(4)
							} else {
								v50 = v43
							}
							v61 = v50
						} else {
							v51 = int32(1)
							if v32 != 0 {
								v61 = int32(base.Ui32(v30)>>(uint(v51)%32)) - v51
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_pg_verifymbstr(m, v26+v33, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v64 != v11 {
								F_pfree(m, v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v68 != v16 {
										F_pfree(m, v16)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v72 < int32(3) {
												return base.I64_extend_i32_u(v26)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v25 == v75 {
													return base.I64_extend_i32_u(v26)
												} else {
													F_pfree(m, v25)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int64(0)
													} else {
														return base.I64_extend_i32_u(v26)
													}
												}
											}
										}
									} else {
										v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v72 < int32(3) {
											return base.I64_extend_i32_u(v26)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v25 == v75 {
												return base.I64_extend_i32_u(v26)
											} else {
												F_pfree(m, v25)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v26)
												}
											}
										}
									}
								}
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v68 != v16 {
									F_pfree(m, v16)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v72 < int32(3) {
											return base.I64_extend_i32_u(v26)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v25 == v75 {
												return base.I64_extend_i32_u(v26)
											} else {
												F_pfree(m, v25)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v26)
												}
											}
										}
									}
								} else {
									v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v72 < int32(3) {
										return base.I64_extend_i32_u(v26)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v25 == v75 {
											return base.I64_extend_i32_u(v26)
										} else {
											F_pfree(m, v25)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v26)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = v2
				v26 = F_decrypt_internal(m, v2, int32(1), v11, v16, int32(0), v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = int32(1)
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					v32 = v30 & v28
					if v32 != 0 {
						v33 = v28
					} else {
						v33 = int32(4)
					}
					if v30 == int32(1) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
						if v40 == int32(18) {
							v43 = int32(16)
						} else {
							v43 = int32(0)
						}
						if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v50 = int32(4)
						} else {
							v50 = v43
						}
						v61 = v50
					} else {
						v51 = int32(1)
						if v32 != 0 {
							v61 = int32(base.Ui32(v30)>>(uint(v51)%32)) - v51
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v26+v33, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v64 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v68 != v16 {
									F_pfree(m, v16)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v72 < int32(3) {
											return base.I64_extend_i32_u(v26)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v25 == v75 {
												return base.I64_extend_i32_u(v26)
											} else {
												F_pfree(m, v25)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v26)
												}
											}
										}
									}
								} else {
									v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v72 < int32(3) {
										return base.I64_extend_i32_u(v26)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v25 == v75 {
											return base.I64_extend_i32_u(v26)
										} else {
											F_pfree(m, v25)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v26)
											}
										}
									}
								}
							}
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v68 != v16 {
								F_pfree(m, v16)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v72 < int32(3) {
										return base.I64_extend_i32_u(v26)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v25 == v75 {
											return base.I64_extend_i32_u(v26)
										} else {
											F_pfree(m, v25)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v26)
											}
										}
									}
								}
							} else {
								v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v72 < int32(3) {
									return base.I64_extend_i32_u(v26)
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v25 == v75 {
										return base.I64_extend_i32_u(v26)
									} else {
										F_pfree(m, v25)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v26)
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
func F_pgp_sym_encrypt_bytea(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn14352(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v11 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)) = uint8(v11)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v3
	if v10 == v3 {
		F_px_debug(m, int32(_a_F_pgp_write_pubenc_sesskey_0), int32(0))
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
