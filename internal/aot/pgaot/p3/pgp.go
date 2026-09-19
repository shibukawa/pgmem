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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v556 int32
	_ = v556
	var v560 int64
	_ = v560
	var v563 int64
	_ = v563
	var v566 int64
	_ = v566
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
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
	return v648
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
	v648 = int32(-13)
	goto L1
L5:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_pushf_free_all(m, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
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
		v639 = v18
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
		v639 = v110
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
		v639 = v123
		goto L5
	} else {
		goto L37
	}
L37:
	;
	goto L11
L38:
	;
	if v349 < int32(0) {
		v639 = v349
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v217
	if v217 == int32(0) {
		goto L39
	} else {
		goto L65
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v154
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
		v639 = int32(-17)
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
	v177 = l0 + int32(132)
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
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
	v349 = v215
	goto L38
L65:
	;
	base.MemoryCopy(m, l0+int32(132), l0+int32(11), v217)
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
	v264 = F_pgp_cfb_create(m, v11+int32(276), v231, l0+int32(11), v261, v262, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	v292 = v253
	goto L71
L71:
	;
	v293 = int32(195)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+276)) = uint8(v293)
	v299 = v292 - (v11 + int32(16))
	if v299 <= int32(191) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v264 < int32(0) {
		v349 = v264
		goto L38
	} else {
		goto L73
	}
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v272 = F_pgp_cfb_encrypt(m, v268, v11+int32(8), int32(1), v253)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v280 = F_pgp_cfb_encrypt(m, v274, l0+int32(132), v277, v253+int32(1))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	F_pgp_cfb_free(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v287 = v285 + int32(1)
	if v287 < int32(0) {
		v349 = v287
		goto L38
	} else {
		goto L77
	}
L77:
	;
	v292 = v287 + v253
	goto L71
L78:
	;
	v336 = F_pushf_write(m, v228, v11+int32(276), v332-(v11+int32(276)))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L85
	}
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v299)
	v332 = v11 + int32(278)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(v299) <= base.Ui32(int32(_a_F_pgp_encrypt_1)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v308 = v299 - int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+278)) = uint8(v308)
	v313 = int32(base.Ui32(v308)>>(uint(int32(8))%32)) + int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v313)
	v332 = v11 + int32(279)
	goto L78
L83:
	;
	goto L84
L84:
	;
	v317 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+277)) = uint8(v317)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+281)) = uint8(v299)
	v321 = int32(base.Ui32(v299) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+280)) = uint8(v321)
	v324 = int32(base.Ui32(v299) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+279)) = uint8(v324)
	v327 = int32(base.Ui32(v299) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+278)) = uint8(v327)
	v332 = v11 + int32(282)
	goto L78
L85:
	;
	if int32(0) <= v336 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v342 = F_pushf_write(m, v228, v11+int32(16), v299)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	v344 = v336
	goto L88
L88:
	;
	if v299 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v344 = v342
	goto L88
L90:
	;
	v349 = v344
	goto L38
L91:
	;
	base.MemoryFill(m, v11+int32(16), int32(0), v299)
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v358 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v359 = int32(-55)
	goto L97
L96:
	;
	v359 = int32(-46)
	goto L97
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v359)
	v364 = F_pushf_write(m, v355, v11+int32(16), int32(1))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	if v364 < int32(0) {
		v639 = v364
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v369 = v11 + int32(8)
	v371 = F_pushf_create(m, v369, int32(_a_F_pgp_encrypt_2), l0, v355)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	if v371 < int32(0) {
		v639 = v371
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v375
	v378 = F_pushf_create(m, v369, int32(_a_F_pgp_encrypt_3), l0, v375)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	if v378 < int32(0) {
		v639 = v378
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v384 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v388 = F_pushf_create(m, v369, int32(_a_F_pgp_encrypt_4), l0, v382)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L107
	}
L105:
	;
	v394 = v382
	goto L106
L106:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v400 = v397 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v400))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v400)%32))&int32(1) == int32(0)) != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if v388 < int32(0) {
		v639 = v388
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v392
	v394 = v392
	goto L106
L109:
	;
	v418 = int32(0)
	v422 = m.G0
	v424 = v422 - int32(16)
	m.G0 = v424
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = v418
	v430 = F_open(m, int32(_a_F_pgp_encrypt_0), v418, v424)
	mBase = m.M
	if v430 != int32(-1) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v417 = int32(0)
	goto L112
L111:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v400<<(uint(int32(2))%32))+uint32(_c_F_pgp_encrypt[0])))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	v417 = v416
	goto L112
L112:
	;
	goto L109
L113:
	;
	if v463 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v433 = int32(1)
	if v417 == int32(0) {
		v456 = v433
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v463 = v418
	goto L116
L116:
	;
	m.G0 = v424 + int32(16)
	goto L113
L117:
	;
	v458 = F_close(m, v430)
	mBase = m.M
	v463 = v456
	goto L116
L118:
	;
	v436 = v11 + int32(16)
	v437 = v417
	goto L119
L119:
	;
	v442 = F_read(m, v430, v436, v437)
	mBase = m.M
	if v442 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v456 = v433
	goto L117
L121:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_pgp_encrypt[1]))
	if v446 == int32(27) {
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v451 = v437 - v442
	if v451 != 0 {
		v436 = v436 + v442
		v437 = v451
		goto L119
	} else {
		goto L125
	}
L124:
	;
	v456 = int32(0)
	goto L117
L125:
	;
	goto L120
L126:
	;
	v639 = int32(-17)
	goto L5
L127:
	;
	goto L128
L128:
	;
	v472 = v11 + int32(16)
	v473 = v472 + v417
	v474 = int32(2)
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473-v474))))
	*(*uint16)(unsafe.Add(mBase, uint32(v473))) = uint16(v476)
	v479 = v417 + v474
	v480 = F_pushf_write(m, v394, v472, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	if v479 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v480 < int32(0) {
		v639 = v480
		goto L5
	} else {
		goto L134
	}
L131:
	;
	base.MemoryFill(m, v472, int32(0), v479)
	goto L133
L132:
	;
	goto L133
L133:
	;
	goto L130
L134:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v486 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v547 != 0 {
		goto L153
	} else {
		goto L154
	}
L136:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v544 = v489
	goto L135
L137:
	;
	goto L138
L138:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v491 <= int32(0) {
		v544 = v490
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v496 = m.G0
	v498 = v496 - int32(16)
	m.G0 = v498
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+14)) = uint8(v500)
	v502 = int32(200)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+15)) = uint8(v502)
	v507 = F_pushf_write(m, v490, v498+int32(15), int32(1))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L141
	}
L140:
	;
	m.G0 = v498 + int32(16)
	if v535 < int32(0) {
		v639 = v535
		goto L5
	} else {
		goto L152
	}
L141:
	;
	if v507 < int32(0) {
		v535 = v507
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v514 = F_pushf_create(m, v498+int32(8), int32(_a_F_pgp_encrypt_2), l0, v490)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v514 < int32(0) {
		v535 = v514
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v522 = F_pushf_write(m, v518, v498+int32(14), int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	if int32(0) <= v522 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v527 = F_pgp_compress_filter(m, v11+int32(8), l0, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	v531 = v522
	goto L148
L148:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	F_pushf_free(m, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L151
	}
L149:
	;
	if int32(0) <= v527 {
		v535 = v527
		goto L140
	} else {
		goto L150
	}
L150:
	;
	v531 = v527
	goto L148
L151:
	;
	v535 = v531
	goto L140
L152:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v541
	v544 = v541
	goto L135
L153:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v550 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v553 = int32(98)
	goto L155
L155:
	;
	v554 = F_time(m)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v554)
	v556 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v556)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v553)
	v560 = int64(base.Ui64(v554) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v560)
	v563 = int64(base.Ui64(v554) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+19)) = uint8(v563)
	v566 = int64(base.Ui64(v554) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v566)
	v568 = int32(203)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+287)) = uint8(v568)
	v573 = F_pushf_write(m, v544, v11+int32(287), int32(1))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L159
	}
L156:
	;
	v551 = int32(117)
	goto L158
L157:
	;
	v551 = int32(116)
	goto L158
L158:
	;
	v553 = v551
	goto L155
L159:
	;
	if v573 < int32(0) {
		v639 = v573
		goto L5
	} else {
		goto L160
	}
L160:
	;
	v580 = F_pushf_create(m, v11+int32(276), int32(_a_F_pgp_encrypt_2), l0, v544)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	if v580 < int32(0) {
		v639 = v580
		goto L5
	} else {
		goto L162
	}
L162:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	v588 = F_pushf_write(m, v584, v11+int32(16), int32(6))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v11)+276))
	if v588 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_pushf_free(m, v590)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v590
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v597 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v639 = v588
	goto L5
L168:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v615 = v613 - v614
	goto L173
L169:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v600 == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v606 = F_pushf_create(m, v11+int32(8), int32(_a_F_pgp_encrypt_5), l0, v590)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v606 < int32(0) {
		v639 = v606
		goto L5
	} else {
		goto L172
	}
L172:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v610
	goto L168
L173:
	;
	v619 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v619)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v622
	v624 = v621 - v622
	if v615 < v624 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v631 = F_pushf_write(m, v629, v630, v626)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L6
	} else {
		goto L178
	}
L175:
	;
	v626 = v615
	goto L177
L176:
	;
	v626 = v624
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v622 + v626
	goto L174
L178:
	;
	if v631 < int32(0) {
		v639 = v631
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v636 = F_pushf_flush(m, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	v639 = v636
	goto L5
L181:
	;
	v648 = v639
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
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
				v34 = F_decrypt_internal(m, v32, v32, v11, v16, v30, v31)
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
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
						if v49 == int32(18) {
							v52 = int32(16)
						} else {
							v52 = int32(0)
						}
						if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v59 = int32(4)
						} else {
							v59 = v52
						}
						v70 = v59
					} else {
						v60 = int32(1)
						if v42 != 0 {
							v70 = int32(base.Ui32(v40)>>(uint(v60)%32)) - v60
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
							v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v43, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v73 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v77 != v16 {
									F_pfree(m, v16)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v81 < int32(3) {
											return v34
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v84 != v30 {
												F_pfree(m, v30)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v89 = v88
													if base.I32_extend16_s(v89) < int32(4) {
														return v34
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v31 == v93 {
															return v34
														} else {
															F_pfree(m, v31)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											} else {
												v89 = v81
												if base.I32_extend16_s(v89) < int32(4) {
													return v34
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v31 == v93 {
														return v34
													} else {
														F_pfree(m, v31)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
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
									v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v81 < int32(3) {
										return v34
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v84 != v30 {
											F_pfree(m, v30)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v89 = v88
												if base.I32_extend16_s(v89) < int32(4) {
													return v34
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v31 == v93 {
														return v34
													} else {
														F_pfree(m, v31)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															return v34
														}
													}
												}
											}
										} else {
											v89 = v81
											if base.I32_extend16_s(v89) < int32(4) {
												return v34
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v31 == v93 {
													return v34
												} else {
													F_pfree(m, v31)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
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
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v77 != v16 {
								F_pfree(m, v16)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v81 < int32(3) {
										return v34
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v84 != v30 {
											F_pfree(m, v30)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
												v89 = v88
												if base.I32_extend16_s(v89) < int32(4) {
													return v34
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v31 == v93 {
														return v34
													} else {
														F_pfree(m, v31)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															return v34
														}
													}
												}
											}
										} else {
											v89 = v81
											if base.I32_extend16_s(v89) < int32(4) {
												return v34
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v31 == v93 {
													return v34
												} else {
													F_pfree(m, v31)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
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
								v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v81 < int32(3) {
									return v34
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v84 != v30 {
										F_pfree(m, v30)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
											v89 = v88
											if base.I32_extend16_s(v89) < int32(4) {
												return v34
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												if v31 == v93 {
													return v34
												} else {
													F_pfree(m, v31)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														return v34
													}
												}
											}
										}
									} else {
										v89 = v81
										if base.I32_extend16_s(v89) < int32(4) {
											return v34
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											if v31 == v93 {
												return v34
											} else {
												F_pfree(m, v31)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
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
						v30 = v22
						v31 = v2
						v32 = int32(1)
						v34 = F_decrypt_internal(m, v32, v32, v11, v16, v30, v31)
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
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
								if v49 == int32(18) {
									v52 = int32(16)
								} else {
									v52 = int32(0)
								}
								if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v59 = int32(4)
								} else {
									v59 = v52
								}
								v70 = v59
							} else {
								v60 = int32(1)
								if v42 != 0 {
									v70 = int32(base.Ui32(v40)>>(uint(v60)%32)) - v60
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							F_pg_verifymbstr(m, v43, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v73 != v11 {
									F_pfree(m, v11)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v77 != v16 {
											F_pfree(m, v16)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v81 < int32(3) {
													return v34
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v84 != v30 {
														F_pfree(m, v30)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v89 = v88
															if base.I32_extend16_s(v89) < int32(4) {
																return v34
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v31 == v93 {
																	return v34
																} else {
																	F_pfree(m, v31)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v89 = v81
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
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
											v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v81 < int32(3) {
												return v34
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v84 != v30 {
													F_pfree(m, v30)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v89 = v88
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v89 = v81
													if base.I32_extend16_s(v89) < int32(4) {
														return v34
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v31 == v93 {
															return v34
														} else {
															F_pfree(m, v31)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
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
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v77 != v16 {
										F_pfree(m, v16)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v81 < int32(3) {
												return v34
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v84 != v30 {
													F_pfree(m, v30)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v89 = v88
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v89 = v81
													if base.I32_extend16_s(v89) < int32(4) {
														return v34
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v31 == v93 {
															return v34
														} else {
															F_pfree(m, v31)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
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
										v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v81 < int32(3) {
											return v34
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v84 != v30 {
												F_pfree(m, v30)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													v89 = v88
													if base.I32_extend16_s(v89) < int32(4) {
														return v34
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v31 == v93 {
															return v34
														} else {
															F_pfree(m, v31)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																return v34
															}
														}
													}
												}
											} else {
												v89 = v81
												if base.I32_extend16_s(v89) < int32(4) {
													return v34
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													if v31 == v93 {
														return v34
													} else {
														F_pfree(m, v31)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
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
							v30 = v22
							v31 = v28
							v32 = int32(1)
							v34 = F_decrypt_internal(m, v32, v32, v11, v16, v30, v31)
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
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
									if v49 == int32(18) {
										v52 = int32(16)
									} else {
										v52 = int32(0)
									}
									if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v59 = int32(4)
									} else {
										v59 = v52
									}
									v70 = v59
								} else {
									v60 = int32(1)
									if v42 != 0 {
										v70 = int32(base.Ui32(v40)>>(uint(v60)%32)) - v60
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
										v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								F_pg_verifymbstr(m, v43, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v73 != v11 {
										F_pfree(m, v11)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v77 != v16 {
												F_pfree(m, v16)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v81 < int32(3) {
														return v34
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v84 != v30 {
															F_pfree(m, v30)
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return int32(0)
															} else {
																v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
																v89 = v88
																if base.I32_extend16_s(v89) < int32(4) {
																	return v34
																} else {
																	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																	if v31 == v93 {
																		return v34
																	} else {
																		F_pfree(m, v31)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int32(0)
																		} else {
																			return v34
																		}
																	}
																}
															}
														} else {
															v89 = v81
															if base.I32_extend16_s(v89) < int32(4) {
																return v34
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v31 == v93 {
																	return v34
																} else {
																	F_pfree(m, v31)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
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
												v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v81 < int32(3) {
													return v34
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v84 != v30 {
														F_pfree(m, v30)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v89 = v88
															if base.I32_extend16_s(v89) < int32(4) {
																return v34
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v31 == v93 {
																	return v34
																} else {
																	F_pfree(m, v31)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v89 = v81
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
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
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v77 != v16 {
											F_pfree(m, v16)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
												if v81 < int32(3) {
													return v34
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v84 != v30 {
														F_pfree(m, v30)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															v89 = v88
															if base.I32_extend16_s(v89) < int32(4) {
																return v34
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
																if v31 == v93 {
																	return v34
																} else {
																	F_pfree(m, v31)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		return v34
																	}
																}
															}
														}
													} else {
														v89 = v81
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
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
											v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v81 < int32(3) {
												return v34
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v84 != v30 {
													F_pfree(m, v30)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
														v89 = v88
														if base.I32_extend16_s(v89) < int32(4) {
															return v34
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
															if v31 == v93 {
																return v34
															} else {
																F_pfree(m, v31)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	return v34
																}
															}
														}
													}
												} else {
													v89 = v81
													if base.I32_extend16_s(v89) < int32(4) {
														return v34
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														if v31 == v93 {
															return v34
														} else {
															F_pfree(m, v31)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13992(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pgp_pub_encrypt_text(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn13992(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
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
	return v798
L2:
	;
	return int32(0)
L3:
	;
	if v24 < int32(0) {
		v798 = v24
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
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	if v778 != 0 {
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
	v768 = v59
	v771 = v48
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
	v768 = v67
	v771 = v48
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
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	F_pullf_free(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L2
	} else {
		goto L187
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
	v721 = v202
	v724 = int32(0)
	v727 = v49
	v728 = v51
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v71
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_0), v20)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L68
	}
L20:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v195 = F_pgp_skip_packet(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	v721 = int32(-114)
	v724 = v74
	v727 = v49
	v728 = int32(1)
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
	v721 = v78
	v724 = v74
	v727 = v49
	v728 = int32(1)
	goto L16
L28:
	;
	v202 = int32(-116)
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
	v721 = v85
	v724 = v87
	v727 = v87
	v728 = v51
	goto L16
L32:
	;
	v202 = int32(-115)
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
		v721 = v99
		v724 = v92
		v727 = v49
		v728 = v51
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
		v721 = v106
		v724 = v92
		v727 = v49
		v728 = v51
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
		v210 = v91
		goto L17
	} else {
		goto L65
	}
L42:
	;
	v721 = int32(-120)
	v724 = v92
	v727 = v49
	v728 = v51
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
		v202 = v119
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
		v202 = v126
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
		v202 = v130
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
	v202 = int32(-103)
	goto L18
L59:
	;
	if v165 < int32(0) {
		v202 = v165
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+103)))
	v173 = F_pgp_cfb_create(m, v20+int32(104), v123, v20+int32(71), v171, int32(0), v164)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v173 < int32(0) {
		v202 = v173
		goto L18
	} else {
		goto L62
	}
L62:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v181 = F_pullf_create(m, v20+int32(108), int32(_a_F_pgp_set_pubkey_2), v180, v91)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	if v181 < int32(0) {
		v202 = v181
		goto L18
	} else {
		goto L64
	}
L64:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v210 = v185
	goto L17
L65:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_3), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v721 = int32(-118)
	v724 = v92
	v727 = v49
	v728 = v51
	goto L16
L67:
	;
	v721 = v195
	v724 = int32(0)
	v727 = v49
	v728 = v51
	goto L16
L68:
	;
	v202 = int32(-107)
	goto L18
L69:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v710 != 0 {
		goto L175
	} else {
		goto L176
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v215
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_4), v20+int32(32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L174
	}
L71:
	;
	v239 = F_pgp_mpi_read(m, v210, v237+v214)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L79
	}
L72:
	;
	v220 = F_pgp_mpi_read(m, v210, v214+int32(24))
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
		v706 = v220
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v226 = F_pgp_mpi_read(m, v210, v214+int32(28))
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
		v706 = v226
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v232 = F_pgp_mpi_read(m, v210, v214+int32(32))
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
		v706 = v232
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
		v706 = v239
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
	if v691 < int32(0) {
		v706 = v691
		goto L69
	} else {
		goto L172
	}
L82:
	;
	v248 = F_pullf_read_fixed(m, v210, int32(20), v20+int32(672))
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
	v327 = F_pullf_read_fixed(m, v210, int32(2), v20+int32(672))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L109
	}
L85:
	;
	if v248 < int32(0) {
		v691 = v248
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
	goto L102
L88:
	;
	if v255 < int32(0) {
		v312 = v255
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
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v20)+656))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v20)+688))
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v20)+640))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v20)+672))
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v20)+648))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v20)+680))
	if base.I64_extend_i32_u(v293^v294)|(v297^v298|(v300^v301)) == int64(0) {
		v312 = v255
		goto L87
	} else {
		goto L99
	}
L99:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_5), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	v312 = int32(-118)
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
	v691 = v312
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
	if v327 < int32(0) {
		v691 = v327
		goto L81
	} else {
		goto L110
	}
L110:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+673)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+672)))
	v337 = int32(0)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
	switch v339 - int32(1) {
	case 0, 1, 2:
		goto L113
	default:
		v683 = v337
		goto L111
	case 15, 16:
		v596 = v337
		v597 = int32(24)
		goto L112
	}
L111:
	;
	if v683 == v331|v332<<(uint(int32(8))%32) {
		v691 = v337
		goto L81
	} else {
		goto L170
	}
L112:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v214+v597)))
	v600 = int32(0)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	v612 = v596 + v606>>(uint(int32(8))%32) + v606&int32(255)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v599)+8))
	if v613 <= v600 {
		v673 = v612
		goto L157
	} else {
		goto L158
	}
L113:
	;
	v343 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v357 = v343 + v351>>(uint(int32(8))%32) + v351&int32(255)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v358 <= v343 {
		v418 = v357
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v214)+28))
	v429 = int32(0)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v441 = v418&int32(_a_F_pgp_set_pubkey_6) + v435>>(uint(int32(8))%32) + v435&int32(255)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	if v442 <= v429 {
		v502 = v441
		goto L129
	} else {
		goto L130
	}
L115:
	;
	goto L114
L116:
	;
	v362 = v358 & int32(3)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if base.Ui32(v358) < base.Ui32(int32(4)) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v402 = v394
	v403 = v395
	v409 = v343
	goto L125
L118:
	;
	v394 = v357
	v395 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v370 = v357
	v371 = int32(0)
	v376 = v343
	goto L121
L121:
	;
	v378 = v371 + v363
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+2)))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+3)))
	v386 = v370 + v379 + v381 + v383 + v385
	v387 = int32(4)
	v388 = v371 + v387
	v390 = v376 + v387
	if v390 != v358&int32(2147483644) {
		v370 = v386
		v371 = v388
		v376 = v390
		goto L121
	} else {
		goto L123
	}
L122:
	;
	if v362 == int32(0) {
		v418 = v386
		goto L115
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v394 = v386
	v395 = v388
	goto L117
L125:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v363))))
	v412 = v402 + v411
	v413 = int32(1)
	v416 = v409 + v413
	if v416 != v362 {
		v402 = v412
		v403 = v403 + v413
		v409 = v416
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v418 = v412
	goto L115
L127:
	;
	goto L126
L128:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v214)+32))
	v513 = int32(0)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v525 = v502&int32(_a_F_pgp_set_pubkey_6) + v519>>(uint(int32(8))%32) + v519&int32(255)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	if v526 <= v513 {
		v586 = v525
		goto L143
	} else {
		goto L144
	}
L129:
	;
	goto L128
L130:
	;
	v446 = v442 & int32(3)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	if base.Ui32(v442) < base.Ui32(int32(4)) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v486 = v478
	v487 = v479
	v493 = v429
	goto L139
L132:
	;
	v478 = v441
	v479 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v454 = v441
	v455 = int32(0)
	v460 = v429
	goto L135
L135:
	;
	v462 = v455 + v447
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+1)))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+2)))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+3)))
	v470 = v454 + v463 + v465 + v467 + v469
	v471 = int32(4)
	v472 = v455 + v471
	v474 = v460 + v471
	if v474 != v442&int32(2147483644) {
		v454 = v470
		v455 = v472
		v460 = v474
		goto L135
	} else {
		goto L137
	}
L136:
	;
	if v446 == int32(0) {
		v502 = v470
		goto L129
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v478 = v470
	v479 = v472
	goto L131
L139:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487+v447))))
	v496 = v486 + v495
	v497 = int32(1)
	v500 = v493 + v497
	if v500 != v446 {
		v486 = v496
		v487 = v487 + v497
		v493 = v500
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v502 = v496
	goto L129
L141:
	;
	goto L140
L142:
	;
	v596 = v586 & int32(_a_F_pgp_set_pubkey_6)
	v597 = int32(36)
	goto L112
L143:
	;
	goto L142
L144:
	;
	v530 = v526 & int32(3)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	if base.Ui32(v526) < base.Ui32(int32(4)) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v570 = v562
	v571 = v563
	v577 = v513
	goto L153
L146:
	;
	v562 = v525
	v563 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v538 = v525
	v539 = int32(0)
	v544 = v513
	goto L149
L149:
	;
	v546 = v539 + v531
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+1)))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+2)))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+3)))
	v554 = v538 + v547 + v549 + v551 + v553
	v555 = int32(4)
	v556 = v539 + v555
	v558 = v544 + v555
	if v558 != v526&int32(2147483644) {
		v538 = v554
		v539 = v556
		v544 = v558
		goto L149
	} else {
		goto L151
	}
L150:
	;
	if v530 == int32(0) {
		v586 = v554
		goto L143
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v562 = v554
	v563 = v556
	goto L145
L153:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571+v531))))
	v580 = v570 + v579
	v581 = int32(1)
	v584 = v577 + v581
	if v584 != v530 {
		v570 = v580
		v571 = v571 + v581
		v577 = v584
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v586 = v580
	goto L143
L155:
	;
	goto L154
L156:
	;
	v683 = v673 & int32(_a_F_pgp_set_pubkey_6)
	goto L111
L157:
	;
	goto L156
L158:
	;
	v617 = v613 & int32(3)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	if base.Ui32(v613) < base.Ui32(int32(4)) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v657 = v649
	v658 = v650
	v664 = v600
	goto L167
L160:
	;
	v649 = v612
	v650 = int32(0)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v625 = v612
	v626 = int32(0)
	v631 = v600
	goto L163
L163:
	;
	v633 = v626 + v618
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+1)))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+2)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+3)))
	v641 = v625 + v634 + v636 + v638 + v640
	v642 = int32(4)
	v643 = v626 + v642
	v645 = v631 + v642
	if v645 != v613&int32(2147483644) {
		v625 = v641
		v626 = v643
		v631 = v645
		goto L163
	} else {
		goto L165
	}
L164:
	;
	if v617 == int32(0) {
		v673 = v641
		goto L157
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	v649 = v641
	v650 = v643
	goto L159
L167:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658+v618))))
	v667 = v657 + v666
	v668 = int32(1)
	v671 = v664 + v668
	if v671 != v617 {
		v657 = v667
		v658 = v658 + v668
		v664 = v671
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v673 = v667
	goto L157
L169:
	;
	goto L168
L170:
	;
	F_px_debug(m, int32(_a_F_pgp_set_pubkey_7), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v691 = int32(-118)
	goto L81
L172:
	;
	v697 = F_pgp_expect_packet_end(m, v210)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	v706 = v697
	goto L69
L174:
	;
	v706 = int32(-118)
	goto L69
L175:
	;
	F_pullf_free(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L2
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v713 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	F_pgp_cfb_free(m, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L2
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v706 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	F_pgp_key_free(m, v214)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L2
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v214
	v721 = v706
	v724 = v214
	v727 = v214
	v728 = v51
	goto L16
L186:
	;
	v721 = v706
	v724 = v92
	v727 = v49
	v728 = v51
	goto L16
L187:
	;
	v733 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v733
	if v724 == v733 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if int32(0) <= v761 {
		v48 = v764
		v49 = v765
		v51 = v728
		goto L5
	} else {
		goto L205
	}
L189:
	;
	v761 = v721
	v763 = v48
	v764 = v48
	v765 = v727
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v721 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v758 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v758
	v761 = v754
	v763 = v756
	v764 = v756
	v765 = v758
	goto L188
L193:
	;
	v754 = v752
	v756 = v48
	goto L192
L194:
	;
	F_pgp_key_free(m, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L2
	} else {
		goto L204
	}
L195:
	;
	v748 = v721
	v749 = v724
	goto L194
L196:
	;
	goto L197
L197:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v724)+48))
	if v739 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v727 == int32(0) {
		v752 = v745
		goto L193
	} else {
		goto L203
	}
L199:
	;
	v745 = v721
	goto L198
L200:
	;
	goto L201
L201:
	;
	if v48 == int32(0) {
		v754 = v721
		v756 = v724
		goto L192
	} else {
		goto L202
	}
L202:
	;
	v745 = int32(-123)
	goto L198
L203:
	;
	v748 = v745
	v749 = v727
	goto L194
L204:
	;
	v752 = v748
	goto L193
L205:
	;
	v768 = v761
	v771 = v763
	goto L7
L206:
	;
	F_pullf_free(m, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L2
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	if v768 < int32(0) {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	goto L208
L210:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L2
	} else {
		goto L219
	}
L211:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	F_pullf_free(m, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L2
	} else {
		goto L218
	}
L212:
	;
	if v771 == int32(0) {
		v788 = v768
		goto L211
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if v771 != 0 {
		goto L210
	} else {
		goto L217
	}
L215:
	;
	F_pgp_key_free(m, v771)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L2
	} else {
		goto L216
	}
L216:
	;
	v788 = v768
	goto L211
L217:
	;
	v788 = int32(-119)
	goto L211
L218:
	;
	v798 = v788
	goto L1
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v771
	v798 = int32(0)
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
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v20 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v24
					v27 = F_decrypt_internal(m, v2, int32(1), v12, v17, int32(0), v26)
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
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
							if v42 == int32(18) {
								v45 = int32(16)
							} else {
								v45 = int32(0)
							}
							if base.Ui32((v42-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v52 = int32(4)
							} else {
								v52 = v45
							}
							v63 = v52
						} else {
							v53 = int32(1)
							if v35 != 0 {
								v63 = int32(base.Ui32(v33)>>(uint(v53)%32)) - v53
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						F_pg_verifymbstr(m, v36, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v66 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v70 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
											if v74 < int32(3) {
												return v27
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v26 == v77 {
													return v27
												} else {
													F_pfree(m, v26)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														return v27
													}
												}
											}
										}
									} else {
										v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v74 < int32(3) {
											return v27
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v26 == v77 {
												return v27
											} else {
												F_pfree(m, v26)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								}
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v70 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v74 < int32(3) {
											return v27
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v26 == v77 {
												return v27
											} else {
												F_pfree(m, v26)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								} else {
									v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v74 < int32(3) {
										return v27
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v26 == v77 {
											return v27
										} else {
											F_pfree(m, v26)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
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
				v26 = v2
				v27 = F_decrypt_internal(m, v2, int32(1), v12, v17, int32(0), v26)
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
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
						if v42 == int32(18) {
							v45 = int32(16)
						} else {
							v45 = int32(0)
						}
						if base.Ui32((v42-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v52 = int32(4)
						} else {
							v52 = v45
						}
						v63 = v52
					} else {
						v53 = int32(1)
						if v35 != 0 {
							v63 = int32(base.Ui32(v33)>>(uint(v53)%32)) - v53
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
							v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					F_pg_verifymbstr(m, v36, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v66 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v70 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v74 < int32(3) {
											return v27
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v26 == v77 {
												return v27
											} else {
												F_pfree(m, v26)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													return v27
												}
											}
										}
									}
								} else {
									v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v74 < int32(3) {
										return v27
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v26 == v77 {
											return v27
										} else {
											F_pfree(m, v26)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												return v27
											}
										}
									}
								}
							}
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v70 != v17 {
								F_pfree(m, v17)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v74 < int32(3) {
										return v27
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v26 == v77 {
											return v27
										} else {
											F_pfree(m, v26)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												return v27
											}
										}
									}
								}
							} else {
								v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v74 < int32(3) {
									return v27
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v26 == v77 {
										return v27
									} else {
										F_pfree(m, v26)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
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
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn13992(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
